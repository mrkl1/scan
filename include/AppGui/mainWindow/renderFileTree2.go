package mainWindow

import (
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/myProj/scaner/new/include/config/settings"
	"github.com/myProj/scaner/new/include/detectOldOfficeExtension"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/myProj/scaner/new/include/searchFilter"
	"github.com/myProj/scaner/new/include/textSearchAndExtract"
	"github.com/myProj/scaner/new/include/unarchive"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
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

func renderFileTree(guiC *appStruct.GuiComponent,btnStart ,btnChooseDir ,btnStop *appStruct.CustomButton) {

	logggerScan.SaveToLog("Start renderFileTree")
	log.Println("START ")
	defer func() { btnStop.SetEnabledFromGoroutine(false)
		btnStart.SetEnabledFromGoroutine(true)
		btnChooseDir.SetEnabledFromGoroutine(true)
	}()

	defer func() {
		guiC.EndUIUpdate   <- ""
		guiC.SearchIsActive = true
		guiC.EndDeleteTemp <- true
	}()


	startDir := guiC.StartDirectoryName
	guiC.FileTree.Clear()

	RemoveAllRows(guiC.ErrorTable)
	RemoveAllRows(guiC.NonScanTable)
	files, err := ioutil.ReadDir(guiC.StartDirectoryName)
	if err != nil {
		guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Невозможно сканировать данную директорию, по причине ее отсутствия или невозможности доступа к ней")
		return
	}

	fileCount := computeFilesCount(startDir, guiC)
	guiC.FileProgress.SetMinimum(0)
	guiC.FileProgress.SetMaximum(fileCount)

	guiC.ProgressBarValue = 0

	guiC.FileProgress.ValueChangedFromGoroutine(guiC.ProgressBarValue)

	go GUIUpdater(guiC)

	for _, file := range files {
		if !guiC.SearchIsActive {
			return
		}

		if file.IsDir() {
			scanDirTree(guiC,file)
		} else {
			scanFileTree(guiC,file)
		}

	}

	guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Сканирование завершено")

}

func scanDirTree(guiC *appStruct.GuiComponent,file os.FileInfo){
	//стартовая директория для поиска
	startFilePath := filepath.Join(guiC.StartDirectoryName, file.Name())
	var head *widgets.QTreeWidgetItem
	err := filepath.Walk(startFilePath,
		func(path string, info os.FileInfo, err error) error {

			defer func(){
				guiC.ProgressBarValue++
			}()

			if !guiC.SearchIsActive {
				return errors.New("break cycle")
			}

			if info == nil {
				return nil
			}

			if info.IsDir(){
				return nil
			}

			ext := detectFileExtension(path)
			guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Сканируется "+path)


			if ext == "" {

				addErrorsToTable(guiC.ErrorTable,
					unarchive.ArchInfoError{
						ArchiveName: path,
						OpenError:   errors.New("Неизвестное расширение"),

					},ext)

				return nil
			}

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
					return nil
				}

				if searchFilter.IsArchive(ext) {
					path = unarchive.CheckExtension(path, ext)
					statArches, errs := unarchive.UnpackWithCtx(path, ext, "", guiC)
					if errs != nil {
						for _, er := range errs {
							addErrorsToTable(guiC.ErrorTable, er, ext)
						}
					}

					if statArches != nil {

						for _, archFile := range statArches {
							stat, containsWord := checkResultFor(archFile.WordFrequency)

							if containsWord {
								if head == nil {
									head = addParent(guiC.FileTree, file, startFilePath)
								}

								rel, _ := filepath.Rel(startFilePath, archFile.Name)
								addChild2(head, newChildNode(stat, rel, archFile.Ext))

							}
						}
					} // конец if проверяющим архивы
				}

				w := textSearchAndExtract.FindText(path,ext,newWordsConfig.GetDictWords(),guiC)

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

func scanFileTree(guiC *appStruct.GuiComponent,file os.FileInfo){
	startFilePath := filepath.Join(guiC.StartDirectoryName, file.Name())
	var head *widgets.QTreeWidgetItem
	ext := detectFileExtension(startFilePath)
	guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine( "Сканируется "+startFilePath+" ")

	logggerScan.SaveToLog(startFilePath+"[ext:"+ext+"]")

	defer func(){
		guiC.ProgressBarValue++
	}()

	if ext == "" {
		addErrorsToTable(guiC.ErrorTable,
			unarchive.ArchInfoError{
				ArchiveName: startFilePath,
				OpenError:   errors.New("Неизвестное расширение"),
			}, ext)
		return
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
			return
		}

		if searchFilter.IsArchive(ext) {
			startFilePath = unarchive.CheckExtension(startFilePath,ext)
			statArches,errs := unarchive.UnpackWithCtx(startFilePath,ext,"",guiC)
			if errs != nil {
				for _,er:= range errs{
					addErrorsToTable(guiC.ErrorTable,er,ext)
				}
			}

			if statArches != nil {
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
		}

		}//isArchive end

		w := textSearchAndExtract.FindText(startFilePath,ext,newWordsConfig.GetDictWords(),guiC)
		stat,containsWord := checkResultFor(w)
		fmt.Println(startFilePath,"textSearchAndExtract ",w)
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
				guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Идет подсчет файлов и папок. На текущий момент обнаружено: "+strconv.Itoa(count))
			}
		}
	}()

	go func(){
		for {
			time.Sleep(50*time.Millisecond)
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


	//https://stackoverflow.com/questions/31888955/check-whether-a-file-is-a-hard-link
	// True if the file is a symlink.
	//https://www.socketloop.com/tutorials/golang-create-and-resolve-read-symbolic-links
	fi,_ := os.Lstat(startDir)
	var newPath string

	if fi.Mode()&os.ModeSymlink != 0 {
		newPath, _ := os.Readlink(startDir)
		//выдается путь вида dir/... а не /dir/...
		//из-за чего путь неправильно распознается
		if runtime.GOOS == "linux" {
			newPath = string(filepath.Separator)+newPath
		}
		newPath,_  = filepath.Abs(newPath)


	} else {
		newPath = startDir
	}

 		filepath.Walk(newPath,
		func(path string, info os.FileInfo, err error) error {
			count++
			if err != nil {
				//println(err,"::: "+path)
			}

			return nil
		})

	stopGorun <- true
	stopGorunSearch <- true
	guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Общее число файлов папок для сканирования: "+strconv.Itoa(count))
	wg.Wait()
	return count
}

func detectFileExtension(path string)string{

	//время по истчении которого будет сделан вывод о том, что
	//невозможно определить расширение файла
	timer := time.NewTimer(1 * time.Second)
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

	 select {
	   case <- res: {

	   }
	}

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
func addParent(tw *appStruct.CustomTreeWidget ,file os.FileInfo,dirPath string)(*widgets.QTreeWidgetItem){
	tv1 := widgets.NewQTreeWidgetItem3(nil,0)
	logggerScan.SaveToLog("head "+dirPath)

	setTreeWidgetItemText(tv1,columnFileName,filepath.Base(dirPath))
	if !file.IsDir(){
		setTreeWidgetItemText(tv1,columnExtensionName,detectFileExtension(dirPath))
	}

	tw.AddTopLevelItemFromGoroutine(tv1)
	logggerScan.SaveToLog("head Text "+tv1.Text(0))
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

	if errs.OpenError.Error() == "Неизвестное расширение" && !settings.IsNeedToShowError(){
		return
	}

	table.SetRowCount(table.RowCount()+1)

	newEntry := widgets.NewQTableWidgetItem2(errs.ArchiveName,0)
	table.SetItem(table.RowCount()-1,0,newEntry)

	newEntry = widgets.NewQTableWidgetItem2(errs.OpenError.Error(),0)
	table.SetItem(table.RowCount()-1,1,newEntry)

	newEntry = widgets.NewQTableWidgetItem2(errs.Ext,0)
	table.SetItem(table.RowCount()-1,2,newEntry)

}




