package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"path/filepath"
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
func newFileTree(guiC *appStruct.GuiComponent)*widgets.QTreeWidget{
	fileTree := widgets.NewQTreeWidget(nil)
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
	fileTree.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	fileTree.ConnectCustomContextMenuRequested(func(pos *core.QPoint){

		menuOpenFile.ConnectTriggered(func(checked bool) {
			si := fileTree.SelectedItems()
			if si != nil && len(si)>0 && si[0].Text(0) != ""  {

				copySi := si[0]
				path := ""
				for copySi != nil {
					path = copySi.Text(0)+"/"+path
					copySi = copySi.Parent()
				}

				path = filepath.Clean(guiC.StartDirectoryName+path)
				//TODO сделать настройку для для архивов
				//     3 колонка определяется расширение
				info := core.NewQFileInfo3(path)
				url := core.QUrl_FromLocalFile(info.AbsoluteFilePath())
				gui.QDesktopServices_OpenUrl(url)

			}

		})

		menuOpenDir.ConnectTriggered(func(checked bool) {
			si := fileTree.SelectedItems()
			if si != nil && len(si)>0 && si[0].Text(0) != ""  {

				copySi := si[0]
				path := ""
				for copySi != nil {
					path = copySi.Text(0)+"/"+path
					copySi = copySi.Parent()
				}

				path = filepath.Clean(guiC.StartDirectoryName+path)

				info := core.NewQFileInfo3(filepath.Dir(path))
				url := core.QUrl_FromLocalFile(info.AbsoluteFilePath())
				gui.QDesktopServices_OpenUrl(url)
			}

		})

		menuOpenBranch.ConnectTriggered(func(checked bool) {
			si := fileTree.SelectedItems()
			if si != nil && len(si)>0 && si[0].Text(0) != ""  {

				expand(si[0])
			}

		})

		ContextMenu.Exec2(fileTree.MapToGlobal(pos), nil)
	})




	return fileTree
}

func resizeColumnTree(fileTree *widgets.QTreeWidget){
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




















