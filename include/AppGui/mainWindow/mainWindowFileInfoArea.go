package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/therecipe/qt/widgets"
)

func newInfoAreaLayout(component *appStruct.GuiComponent)*widgets.QVBoxLayout{
	infoLayout    :=  widgets.NewQVBoxLayout()
	infoTab := widgets.NewQTabWidget(nil)
	infoTab.AddTab(component.FileTree, "Найденные файлы")
	infoTab.AddTab(component.ErrorTable, "Ошибки")
	infoTab.AddTab(component.NonScanTable, "Медиа файлы")
	infoLayout.Layout().AddWidget(infoTab)
	return infoLayout
}


