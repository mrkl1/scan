package mainMenu

import (
	"github.com/myProj/scaner/new/include/config/extensions"
	"github.com/myProj/scaner/new/include/config/settings"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
	sT,checkBoxesSet := settingsTab(exclWidget)
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

func settingsTab(exclWidget *widgets.QWidget)(*widgets.QWidget,[]*widgets.QCheckBox){
	settingsList := widgets.NewQWidget(nil,0)
	vbox    := widgets.NewQVBoxLayout()
	//vbox.SetContentsMargins(0,0,0,0)
	//vbox.SetSpacing(0)

	//vbox.SetAlignment2(vbox,core.Qt__AlignTop)
	settingsList.SetLayout(vbox)
	sets := settings.ReadSettingsFromConfig()
	spinLimit := limitArchSetting()
	spinName := widgets.NewQLabel2("Ограничение на минимальное свободное место при распаковке архивов в ГБ",spinLimit,0)



	vbox.AddWidget(spinName,0,0)
	vbox.AddWidget(spinLimit,0,0)




	var checkBoxes []*widgets.QCheckBox

	//m := newSetMap()
	for _,set := range sets {

		if set.Setting == "Ограничение на минимальное свободное место при распаковке архивов в ГБ"{
			continue
		}
		chbx := widgets.NewQCheckBox2(set.Setting,nil)
		chbx.SetChecked(false)
		if set.IsAllowSetting {
			chbx.SetChecked(true)
		}

		checkBoxes = append(checkBoxes,chbx)

		vbox.AddWidget(chbx,0,0)
		//i,j = getNewPosition(i,j,3)
	}

	exclWidget.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		settings.SetArchiveLimit(spinLimit.Value())
		sets := settings.ReadSettingsFromConfig()
		for i,c := range checkBoxes {

			if c.Text() == "Отображать файлы в таблице с неизвестным расширением"{
				sets[i].IsAllowSetting = c.IsChecked()

			}
		}

		settings.SetNewConfig(sets)

	})

	vbox.AddStretch(0)

	return settingsList,checkBoxes
}

func limitArchSetting() *widgets.QSpinBox{
	spinBox := widgets.NewQSpinBox(nil)
	spinBox.SetMinimum(0)
	spinBox.SetMaximum(1000)
	spinBox.SetValue(settings.GetArchiveLimit())
	return spinBox
}


func newSetMap()map[string]string{
	m := make(map[string]string)
	m["allowUnknownExtension"] = "Отображать файлы в таблице с неизвестным расширением"
	//m["restrictCountOfLinesInTable"] = "Ограничить число строк в таблиц"
	//m["restrictCountOfLinesInTable"] = "Записывать неизвестные ошибки в файл"
	return m
}

