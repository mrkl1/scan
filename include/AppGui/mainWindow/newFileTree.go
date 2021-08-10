package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/myProj/scaner/new/include/searchFilter"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	columnCount = 3
	columnFileName = "Файлы"
	columnStatisticsName = "Частота вхождения слов"
	columnExtensionName = "Тип файла"
	columnpath = "path"
)
var colNames = []string{columnFileName,columnStatisticsName,columnExtensionName,columnpath}


//newFileTree дерево для отображения найденных файлов//newFileTree дерево для отображения найденных файлов//newFileTree дерево для отображения найденных файлов//newFileTree дерево для отображения найденных файлов
func newFileTree(guiC *appStruct.GuiComponent)*appStruct.CustomTreeWidget{
	fileTree := appStruct.NewCustomTreeWidget(guiC.MainWindow)


	fileTree.SetColumnCount(columnCount)
	fileTree.SetHeaderLabels(colNames)
	fileTree.HideColumn(3)
	//разворачивает дерево при клике на него

	fileTree.ConnectDoubleClicked(func(index *core.QModelIndex){
		resizeColumnTree(fileTree)
		//expand(fileTree.ItemFromIndex(index))
	})

	//cont menu
	ContextMenu := widgets.NewQMenu(fileTree)
	menuOpenFile := ContextMenu.AddAction("Открыть файл")
	menuOpenDir := ContextMenu.AddAction("Открыть директорию файла")
	menuOpenBranch := ContextMenu.AddAction("Раскрыть ветку")


	/*
	чтобы настроить контекстное меню для конкретного элемента
	нужно сделать нужно поставить SetContextMenuPolicy
	и ConnectCustomContextMenuRequested
	 */
	fileTree.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	fileTree.ConnectCustomContextMenuRequested(func(pos *core.QPoint){

		ContextMenu.Exec2(guiC.MainWindow.MapToGlobal(pos), nil)

	})


	menuOpenFile.ConnectTriggered(func(checked bool) {



		si := fileTree.SelectedItems()
		if si != nil && len(si)>0 && si[0].Text(0) != ""  {

			copySi := si[0]
			path := ""
			for copySi != nil {
				path =  filepath.Join(copySi.Text(0),path)
				copySi = copySi.Parent()
			}

			path = filepath.Clean(filepath.Join( guiC.StartDirectoryName,path))
			archPath := checkForArchive(path)
			if  archPath != ""{
				archPath,_ := filepath.Abs(archPath)
				path = archPath
			}

			//TODO чтобы возможно было открыть с подсветкой файла
			// в директории нужно использовать местные файловые менеджеры
			// linux nautilus
			// windows explorer
			//"nautilus","/media/us/Transcend/тестовые_данные_для _программ/dict.txt"
			//cmd := exec.Command("nautilus",archPath)
			//fmt.Println(archPath)
			//err := cmd.Run()
			//if err != nil {
			//}

			info := core.NewQFileInfo3(path)
			url := core.QUrl_FromLocalFile(info.AbsoluteFilePath())
			ok := gui.QDesktopServices_OpenUrl(url)
			if !ok {
				logggerScan.SaveToLog("Не получилось открыть файл/папку:"+path)
			}


		}


	})

	menuOpenDir.ConnectTriggered(func(checked bool) {

		si := fileTree.SelectedItems()
		if si != nil && len(si)>0 && si[0].Text(0) != ""  {

			copySi := si[0]
			path := ""
			for copySi != nil {
				path =  filepath.Join(copySi.Text(0),path)
				copySi = copySi.Parent()
			}

			path = filepath.Clean(filepath.Join( guiC.StartDirectoryName,path))
			archPath := checkForArchive(path)
			if  archPath != ""{
				archPath,_ := filepath.Abs(archPath)
				path = archPath
			}
			info := core.NewQFileInfo3(path)
			url := core.QUrl_FromLocalFile(info.AbsolutePath())
			ok := gui.QDesktopServices_OpenUrl(url)
			if !ok {
				logggerScan.SaveToLog("Не получилось открыть файл/папку:"+path)
			}


		}



	})

	menuOpenBranch.ConnectTriggered(func(checked bool) {
		si := fileTree.SelectedItems()
		if si != nil && len(si)>0 && si[0].Text(0) != ""  {

			expand(si[0])
		}

	})

	return fileTree
}

func resizeColumnTree(fileTree *appStruct.CustomTreeWidget){
	for i := 0;i < fileTree.ColumnCount();i++{
		fileTree.ResizeColumnToContents(i)
	}
}

func expand(ptr *widgets.QTreeWidgetItem){
	for ptr.ChildCount() != 0 {
		ptr.SetExpanded(true)
		ptr = ptr.Child(0)
	}
}


func checkForArchive(path string)string{

	spldr :=splitDirForMenu(path)

	for i,d := range spldr{
		if searchFilter.IsContainArchiveExtension(d){

			newPath := filepath.Join(spldr[:i]...)

			if runtime.GOOS == "linux" {
				newPath = string(filepath.Separator)+newPath
			}
			logggerScan.SaveToLog("ARCH PATH FOR OPEN:"+newPath)
			return newPath
		}
	}
	return ""
}

func splitDirForMenu(dir string)[]string{
	var s []string

	if runtime.GOOS == "windows"{
		winSplit := strings.Split(dir,string(filepath.Separator))
		if len(winSplit)>0 {
			s = append(s,winSplit[0]+`\\`)
			winSplit = winSplit[1:]
		}
		for i,v := range winSplit {
			if v == "" {
				s = append(s[:i],s[i+1:]...)
			}
		}
		logggerScan.SaveToLog("splitDirForMenu(windows)"+strings.Join(s,"") )
		return s
	}


	return splitDir(dir)
}

















