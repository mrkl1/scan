package mainMenu

import (
	"bufio"
	"github.com/gabriel-vasile/mimetype"
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config/extensions"
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/myProj/scaner/new/include/config/words"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/myProj/scaner/new/include/reportScan"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
	"strings"
)

//названия пунктов
const (
	fileMenuName = "&Файл"
	exportName   = "Экспорт списка слов"
	importName   = "Импорт списка слов"
	saveReportName   = "Сохранить отчет"
	fileActionWordWindowName  = "Посмотреть список слов"
	fileActionExitName  = "Выход"
	settingMenuName = "Настройки"
)


//NewBar отрисовывает главное меню, также
//задается обрабокта для всех пунктов в этом меню
func NewBar(component *appStruct.GuiComponent)*widgets.QMenuBar{

	newMenuBar   := widgets.NewQMenuBar(nil)

	fileMenu     := newMenuBar.AddMenu2(fileMenuName)
	fileExport   := fileMenu.AddAction(exportName)
	fileImport     := fileMenu.AddAction(importName)
	saveReport := fileMenu.AddAction(saveReportName)
	settingsMenu := newMenuBar.AddMenu2(settingMenuName)

	settingsExt := settingsMenu.AddAction("Фильтр форматов")
	btnShowWordListWindow :=  fileMenu.AddAction(fileActionWordWindowName)
	btnExit := fileMenu.AddAction(fileActionExitName)

	wordListWindow := newWordListWindow(component)

	//открывает окно с со списком слов для поиска
	btnShowWordListWindow.ConnectTriggered(func(bool) {
		wordListWindow.Show()
	})

	//закрыть приложение
	btnExit.ConnectTriggered(func(bool){
		component.Application.Exit(0)
	})

	saveReport.ConnectTriggered(func(bool){
		reportScan.SaveReport(component)
	})


	fileImport.ConnectTriggered(func(bool) {
		importFile()


	})

	fileExport.ConnectTriggered(func(bool) {
		newWordsConfig.WordExport()
		//updateWordList(component.WordList)
	})

	settingsExt.ConnectTriggered(func(bool){
		ecxludeWindow,checkBoxes  := NewExtensionTab()
		ecxludeWindow.Show()

		ecxludeWindow.ConnectCloseEvent(func(event *gui.QCloseEvent){
			var exts []extensions.Extension
			//var sets []settings.Settings
			for _,c := range checkBoxes.ExtcheckBoxes {
				exts = append(exts,extensions.Extension{
					Ext:          c.Text(),
					AllowScanning: c.IsChecked(),
				})


			}


			err := extensions.SetExtStatus(exts)


			if err != nil {
				logggerScan.SaveToLog("SetExtStatus error:"+err.Error())
			}

			ecxludeWindow.Close()

		})

	})

	return newMenuBar
}

func wordImport(){
	filename := widgets.QFileDialog_GetOpenFileName(nil,"Open Directory","","","",widgets.QFileDialog__DontUseNativeDialog)
	m,_ := mimetype.DetectFile(filename)
	if m.Extension() ==".txt"{
		f, _ := os.Open(filename)
		defer f.Close()
		r := bufio.NewReader(f)
		line, err := r.ReadString('\n')
		for err == nil {
			line = strings.TrimSpace(line)
			i := strings.Index(line," ")
			if i > -1 {
				line = line[:i]
			}
			words.AddWordToConfig(line)
			line, err = r.ReadString('\n')
		}
	}

	if m.Extension() ==".json"{
		words.ImportJson(filename)
	}

}









