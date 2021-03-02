package mainMenu

import (

	"github.com/myProj/scaner/new/include/config/extensions"
	"github.com/myProj/scaner/new/include/config/settings"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type TabCheckBoxSettings struct {
	ExtcheckBoxes []*widgets.QCheckBox
	SetcheckBoxes []*widgets.QCheckBox
}

func NewExtensionTab()(*widgets.QWidget,*TabCheckBoxSettings){
	exclWidget := widgets.NewQWidget(nil,1)
	vbox := widgets.NewQVBoxLayout()
	exclWidget.SetLayout(vbox)
	exclWidget.SetWindowFlags(core.Qt__Dialog)
	infoTab := widgets.NewQTabWidget(nil)
	eT,checkBoxesExt := extensionTab()
	sT,checkBoxesSet := settingsTab()
	infoTab.AddTab(eT, "Список расширений")
	infoTab.AddTab(sT, "Дополнительные настройки")
	vbox.Layout().AddWidget(infoTab)

	TabCheckBoxS := &TabCheckBoxSettings{
		ExtcheckBoxes: checkBoxesExt,
		SetcheckBoxes: checkBoxesSet,
	}

	return exclWidget,TabCheckBoxS
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

func settingsTab()(*widgets.QWidget,[]*widgets.QCheckBox){
	settingsList := widgets.NewQWidget(nil,1)
	vbox    := widgets.NewQGridLayout(nil)
	//
	settingsList.SetLayout(vbox)

	sets := settings.ReadSettingsFromConfig()


	var checkBoxes []*widgets.QCheckBox
	var i,j int
	//m := newSetMap()
	for _,set := range sets {

		chbx := widgets.NewQCheckBox2(set.Setting,nil)
		chbx.SetChecked(false)
		if set.IsAllowSetting {
			chbx.SetChecked(true)
		}

		checkBoxes = append(checkBoxes,chbx)

		vbox.AddWidget2(chbx,i,j,0)
		i,j = getNewPosition(i,j,3)
	}

	return settingsList,checkBoxes
}


func newSetMap()map[string]string{
	m := make(map[string]string)
	m["allowUnknownExtension"] = "Отображать файлы в таблице с неизвестным расширением"
	//m["restrictCountOfLinesInTable"] = "Ограничить число строк в таблиц"
	//m["restrictCountOfLinesInTable"] = "Записывать неизвестные ошибки в файл"
	return m
}

