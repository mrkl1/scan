package mainWindow

import (
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/myProj/scaner/new/include/detectOldOfficeExtension"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/myProj/scaner/new/include/searchFilter"
	"github.com/myProj/scaner/new/include/textSearchAndExtract"
	"github.com/myProj/scaner/new/include/unarchive"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type treeNodeParametrs struct {
	stat       string
	spitedPath []string
	ext        string
}

func renderFileTree(guiC *appStruct.GuiComponent,btnStart *widgets.QPushButton,btnChooseDir *widgets.QPushButton,btnStop *widgets.QPushButton) {
	defer func() { btnStop.SetEnabled(false) }()
	defer func() { guiC.SearchIsActive = true }()
	startDir := guiC.StartDirectoryName
	guiC.FileTree.Clear()
	//RemoveAllRows(guiC.NonScanTable)
	//RemoveAllRows(guiC.ErrorTable)

	fileCount := computeFilesCount(startDir, guiC)
	guiC.FileProgress.SetMinimum(0)
	guiC.FileProgress.SetMaximum(fileCount)
	count := 0
	files, _ := ioutil.ReadDir(guiC.StartDirectoryName)

	for _, file := range files {
		if !guiC.SearchIsActive {
			return
		}

		if file.IsDir() {
			scanDirTree(guiC,&count,file)
		} else {
			scanFileTree(guiC,&count,file)
		}

	}
	guiC.FileProgressUpdate <- count

	guiC.InfoAboutScanningFilesUpdate <- "Сканирование завершено"
	guiC.SearchIsActive = false

	btnStart.SetEnabled(true)
	btnChooseDir.SetEnabled(true)
}

func scanDirTree(guiC *appStruct.GuiComponent,count *int,file os.FileInfo){
	//стартовая директория для поиска
	startFilePath := filepath.Join(guiC.StartDirectoryName, file.Name())
	var head *widgets.QTreeWidgetItem
	err := filepath.Walk(startFilePath,
		func(path string, info os.FileInfo, err error) error {
			if !guiC.SearchIsActive {
				return errors.New("break cycle")
			}

			ext := detectFileExtension(path)
			guiC.InfoAboutScanningFilesUpdate <- "Сканируется "+path

			if ext == ""{
				addErrorsToTable(guiC.ErrorTable,
					unarchive.ArchInfoError{
						ArchiveName: path,
						OpenError:   errors.New("Неизвестное расширение"),

					},ext)
			}

			*count++
			guiC.FileProgressUpdate <- *count


			fmt.Println(ext,path)
			if searchFilter.IsExtensionForSearch(ext) {

				if searchFilter.IsMedia(ext){
					addNonScanItem(guiC.NonScanTable,path,ext)
				}

				if searchFilter.IsUnsupported(ext){
					addErrorsToTable(guiC.ErrorTable,
						unarchive.ArchInfoError{
							ArchiveName: path,
							OpenError:   errors.New("Расширение не поддерживается"),

						},ext)
				}

				if searchFilter.IsArchive(ext){
					path = unarchive.CheckExtension(path,ext)
					statArches,errs := unarchive.UnpackWithCtx(path,ext,"",guiC)
					if errs != nil {
						for _,er:= range errs{
							addErrorsToTable(guiC.ErrorTable,er,ext)
						}
					}
					for _,archFile := range statArches {
						stat,containsWord := checkResultFor(archFile.WordFrequency)
						if containsWord {
							if head == nil {
								head = addParent(guiC.FileTree,file,startFilePath)
							}
							rel,_ := filepath.Rel(startFilePath,archFile.Name)
							addChild2(head,newChildNode(stat,rel ,archFile.Ext))

						}
					}
				}// конец if проверяющим архивы

				w := textSearchAndExtract.FindText(path,ext,newWordsConfig.GetDictWords())

				stat,containsWord := checkResultFor(w)
				if containsWord {
					if head == nil {
						head = addParent(guiC.FileTree,file,startFilePath)
					}
					relPath,_ := filepath.Rel(startFilePath,path)
					addChild2(head,newChildNode(stat,relPath,ext))
				}
			}// конец if проверяющего файлы с подходящим расширением

			return nil
		})

	if err != nil {
		logggerScan.SaveToLog("Walk error"+err.Error())
	}
}

func scanFileTree(guiC *appStruct.GuiComponent,count *int,file os.FileInfo){
	startFilePath := filepath.Join(guiC.StartDirectoryName, file.Name())
	var head *widgets.QTreeWidgetItem
	ext := detectFileExtension(startFilePath)
	guiC.InfoAboutScanningFilesUpdate <- "Сканируется "+startFilePath+" "
	*count++
	guiC.FileProgressUpdate <- *count
	//guiC.FileProgress.SetValue(*count)

	fmt.Println(ext,startFilePath)
	if ext == "" {
		addErrorsToTable(guiC.ErrorTable,
			unarchive.ArchInfoError{
				ArchiveName: startFilePath,
				OpenError:   errors.New("Неизвестное расширение"),
			}, ext)
	}

	if searchFilter.IsExtensionForSearch(ext) {

		if searchFilter.IsMedia(ext){
			addNonScanItem(guiC.NonScanTable,startFilePath,ext)
		}

		if searchFilter.IsUnsupported(ext){
			addErrorsToTable(guiC.ErrorTable,
				unarchive.ArchInfoError{
					ArchiveName: startFilePath,
					OpenError:   errors.New("Расширение не поддерживается"),

				},ext)
		}

		if searchFilter.IsArchive(ext) {
			startFilePath = unarchive.CheckExtension(startFilePath,ext)
			statArches,errs := unarchive.UnpackWithCtx(startFilePath,ext,"",guiC)
			if errs != nil {
				for _,er:= range errs{
					addErrorsToTable(guiC.ErrorTable,er,ext)
				}
			}

			for _,archFile := range statArches {
				stat,containsWord := checkResultFor(archFile.WordFrequency)

				if containsWord {
					if head == nil {
						head = addParent(guiC.FileTree,file,startFilePath)
					}
					rel,_ := filepath.Rel(startFilePath,archFile.Name)
					tnp := newChildNode(stat,rel,archFile.Ext)
					addChild2(head,tnp)

				}
			}

			return
		}//isArchive end
		w := textSearchAndExtract.FindText(startFilePath,ext,newWordsConfig.GetDictWords())
		stat,containsWord := checkResultFor(w)
		if containsWord {
			head = addParent(guiC.FileTree,file,startFilePath)
			setTreeWidgetItemText(head,columnStatisticsName,stat)
			g := gui.NewQBrush3(gui.NewQColor3(0xd3, 0xd3, 0xd3, 0xff), 1)
			head.SetBackground(0,g)
		}
	}

}


//нужно обработать ошибки
func computeFilesCount(startDir string,guiC *appStruct.GuiComponent)int  {
	//чтобы не считало саму директорию
	count := -1
	stopGorun := make(chan bool,1)
	stopGorunSearch := make(chan bool,1)
	defer close(stopGorun)
	defer close(stopGorunSearch)

	//обновляет инфу о найденных файлах
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(){
		for {
			select {
			case <-stopGorun:
				//нужно чтобы сработало событие завершения на кнопке отмены
				guiC.SearchIsActive = true
				wg.Done()
				return
			default:
				//без sleep крашится
				//через 3-4 прогона
				time.Sleep(300*time.Millisecond)
				guiC.InfoAboutScanningFilesUpdate <- "Идет подсчет файлов и папок. На текущий момент обнаружено: "+strconv.Itoa(count)
			}
		}
	}()

	go func(){
		for {
			time.Sleep(10*time.Millisecond)
			if !guiC.SearchIsActive {
				stopGorun <- true
				wg.Done()
				return
			}

			select {
			case <-stopGorunSearch:
				wg.Done()
				return
			default :
				continue
			}


		}
	}()

	filepath.Walk(startDir,
		func(path string, info os.FileInfo, err error) error {
			count++
			if err != nil {
				//println(err,"::: "+path)
			}

			return nil
		})

	stopGorun <- true
	stopGorunSearch <- true
	guiC.InfoAboutScanningFilesUpdate <- "Общее число файлов папок для сканирования: "+strconv.Itoa(count)
	wg.Wait()
	return count
}

func detectFileExtension(path string)string{

	//время по истчении которого будет сделан вывод о том, что
	//невозможно определить расширение файла
	timer := time.NewTimer(2 * time.Second)
	var result = ""
	res := make(chan string)
	go detecting(path,res)

	select {
	//если 2 секунды проходит
	case <-timer.C:
		return ""
	//если нормально определил результат
	case result = <- res:
		// освобождет ресурс
		return result
	}
}

func detecting(path string,res chan string){
	//test на офис
	//проверить на ole
	ext := detectOldOfficeExtension.DetectOldOffice(path)
	if ext != ""{
		res <- ext
	}

	mime, _ := mimetype.DetectFile(path)
	res <- mime.Extension()
}

func addNonScanItem(table *widgets.QTableWidget,path,ext string){
	table.SetRowCount(table.RowCount()+1)
	newEntry := widgets.NewQTableWidgetItem2(path,0)
	table.SetItem(table.RowCount()-1,0,newEntry)
	newEntry = widgets.NewQTableWidgetItem2(ext,0)
	table.SetItem(table.RowCount()-1,1,newEntry)



}

func checkResultFor(wf map[string]int)(output string,isContainWords bool){
	for k,v := range wf {
		if v > 0 {
			if isContainWords{
				output += "\n"+k+": "+fmt.Sprintf("%d",v)
			} else {
				output += k+": "+fmt.Sprintf("%d",v)
			}

			isContainWords = true
		}
	}
	return output,isContainWords
}

func setTreeWidgetItemText(item *widgets.QTreeWidgetItem,columnName,text string){
	for i,v := range colNames{
		if v == columnName{
			item.SetText(i,text)
			return
		}
	}
}


//в конфиге нужно определить
//кол-во колонок
func addParent(tw *widgets.QTreeWidget ,file os.FileInfo,dirPath string)(*widgets.QTreeWidgetItem){
	tv1 := widgets.NewQTreeWidgetItem3(nil,0)

	setTreeWidgetItemText(tv1,columnFileName,filepath.Base(dirPath))
	if !file.IsDir(){
		setTreeWidgetItemText(tv1,columnExtensionName,detectFileExtension(dirPath))
	}
	tw.AddTopLevelItem(tv1)
	return tv1
}

func splitDir(dir string)[]string{
	s := strings.Split(dir,string(filepath.Separator))
	for i,v := range s {
		if v == "" {
			s = append(s[:i],s[i+1:]...)
		}
	}
	return s
}

func newChildNode(stat,relPath,ext string)treeNodeParametrs{
	tnp := treeNodeParametrs{
		stat:       stat,
		spitedPath: splitDir(relPath),
		ext:        ext,
	}
	return tnp
}

func addChild2(parent *widgets.QTreeWidgetItem,  newElements treeNodeParametrs)*widgets.QTreeWidgetItem{

	for _,newEl := range newElements.spitedPath{
		if newEl =="."{
			continue
		}
		var child *widgets.QTreeWidgetItem
		//проверка на то есть ли этот
		//элемент уже в дереве
		var exist = false
		var ind = 0
		for i := 0;i < parent.ChildCount();i++ {
			if parent.Child(i).Text(0) == newEl{
				ind = i
				exist = true
			}
		}
		if !exist{
			child = widgets.NewQTreeWidgetItem(0)
			setTreeWidgetItemText(child,columnFileName,newEl)
			parent.AddChild(child)
			parent = child
		} else {
			parent = parent.Child(ind)
		}
	}

	setTreeWidgetItemText(parent,columnStatisticsName,newElements.stat)
	setTreeWidgetItemText(parent,columnExtensionName,newElements.ext)

	//set for files color
	g := gui.NewQBrush3(gui.NewQColor3(0xd3, 0xd3, 0xd3, 0xff), 1)
	parent.SetBackground(0,g)
	parent = getTopLevelItem(parent)
	return parent
}

func getTopLevelItem(parent *widgets.QTreeWidgetItem) *widgets.QTreeWidgetItem{
	for {
		if parent.Parent().Parent() == nil {
			return parent
		} else {
			parent = parent.Parent()
		}
	}
}

func addErrorsToTable(table *widgets.QTableWidget,errs unarchive.ArchInfoError,ext string){

	table.SetRowCount(table.RowCount()+1)

	newEntry := widgets.NewQTableWidgetItem2(errs.ArchiveName,0)
	table.SetItem(table.RowCount()-1,0,newEntry)

	newEntry = widgets.NewQTableWidgetItem2(errs.OpenError.Error(),0)
	table.SetItem(table.RowCount()-1,1,newEntry)

	newEntry = widgets.NewQTableWidgetItem2(ext,0)
	table.SetItem(table.RowCount()-1,2,newEntry)

}




