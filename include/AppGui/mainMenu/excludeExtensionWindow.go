package mainMenu

import (
	"github.com/myProj/scaner/new/include/config/extensions"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func NewExtensionTab()(*widgets.QWidget,[]*widgets.QCheckBox){
	exclWidget := widgets.NewQWidget(nil,1)
	vbox := widgets.NewQVBoxLayout()
	exclWidget.SetLayout(vbox)
	exclWidget.SetWindowFlags(core.Qt__Dialog)
	infoTab := widgets.NewQTabWidget(nil)
	eT,checkBoxes := extensionTab()
	infoTab.AddTab(eT, "Список расширений")
	//infoTab.AddTab(widgets.NewQWidget(nil,1), "Ошибки")

	vbox.Layout().AddWidget(infoTab)
	return exclWidget,checkBoxes
}

func extensionTab()(*widgets.QWidget,[]*widgets.QCheckBox){
	extList := widgets.NewQWidget(nil,1)
	vbox    := widgets.NewQGridLayout(nil)
	//
	extList.SetLayout(vbox)
	exts := extensions.GetAllowList()


	var checkBoxes []*widgets.QCheckBox
	var i,j int
	for _,ext := range exts {

		chbx := widgets.NewQCheckBox2(ext.Ext,nil)
		chbx.SetChecked(false)
		if ext.AllowScanning {
			chbx.SetChecked(true)
		}

		checkBoxes = append(checkBoxes,chbx)

		vbox.AddWidget2(chbx,i,j,0)
		i,j = getNewPosition(i,j,3)
	}

	return extList,checkBoxes
}


func getNewPosition(row,column,limit int)(int,int){
	if column < limit {
		return row , column+1
	}
	return row+1, 0
}