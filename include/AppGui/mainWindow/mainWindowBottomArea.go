package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/myProj/scaner/new/include/tempDeleter"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"strings"
	"time"
)

const(
	btnChooseDirName = "Выбрать директорию"
	btnStartName = "Начать поиск"
	btnStopName = "Прекратить поиск"
	btnSkipName = "Пропустить файл"
)

//переименовать в controlButtonArea
func newBottomArea(guiC *appStruct.GuiComponent)*widgets.QHBoxLayout{
	hbox       := widgets.NewQHBoxLayout()
	hbox.AddLayout(controlButtonsArea(guiC),0)
	hbox.AddLayout(optionalInfo(guiC),0)
	return hbox
}

func controlButtonsArea(guiC *appStruct.GuiComponent)*widgets.QVBoxLayout{
	controlLayout := widgets.NewQVBoxLayout()
	btnChooseDir  := appStruct.NewCustomButton2(btnChooseDirName,nil)
	btnStart      := appStruct.NewCustomButton2(btnStartName,nil)
	btnStop       := appStruct.NewCustomButton2(btnStopName,nil)
	btnSkip       := appStruct.NewCustomButton2(btnSkipName,nil)
	controlLayout.Layout().AddWidget(btnChooseDir)
	controlLayout.Layout().AddWidget(btnSkip)
	controlLayout.Layout().AddWidget(btnStart)
	controlLayout.Layout().AddWidget(btnStop)
	btnStop.SetEnabled(false)
	btnStart.ConnectClicked(func(checked bool) {

		_,err := ioutil.ReadDir(guiC.StartDirectoryName)
		if err != nil {
			guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Невозможно сканировать данную директорию, по причине ее отсутствия или невозможности доступа к ней")
			return
		}

		guiC.SearchIsActive = true
		btnStart.SetEnabled(false)
		btnChooseDir.SetEnabled(false)
		btnStop.SetEnabled(true)

		go tempDeleter.StartDelete(guiC)
		go renderFileTree(guiC,btnStart,btnChooseDir,btnStop)

	})

	btnSkip.ConnectClicked(func(checked bool){
		guiC.SkipItem = true
		guiC.SkipItemNonArch = true
	})

	btnChooseDir.ConnectClicked(func(checked bool){
		filePath := widgets.QFileDialog_GetExistingDirectory(guiC.MainWindow,"Open Directory",guiC.StartDirectoryName,widgets.QFileDialog__DontUseNativeDialog)
		if filePath == "" {
			return
		}

		guiC.StartDirectoryForScan.UpdateTextFromGoroutine(directoryForScanning+": "+filePath)
		guiC.StartDirectoryForScan.AdjustSizeFromGoroutine()
		guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("")
		guiC.StartDirectoryName = filePath
		cfg := config.GetCurrentConfig()
		cfg.StartCoordinate =[]int{
			guiC.MainWindow.Geometry().X(),
			guiC.MainWindow.Geometry().Y(),
			guiC.MainWindow.Width(),
			guiC.MainWindow.Height(),
		}
		cfg.StartDir = guiC.StartDirectoryName

		err := config.SaveCurrentConfig(cfg)
		if err != nil {
			logggerScan.SaveToLog("SaveCurrentConfig error:"+err.Error())
		}
	})

	btnStop.ConnectClicked(func(checked bool) {
		guiC.SearchIsActive = false
		btnStop.SetEnabled(false)
		go func (){
			guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Подождите идет завершение процесса сканирования")
			for {
				time.Sleep(time.Millisecond*350)
				guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine(updateStopInfo(guiC.InfoAboutScanningFiles))
				if guiC.SearchIsActive{
					guiC.InfoAboutScanningFiles.UpdateTextFromGoroutine("Процесс сканирования был остановлен")
					btnStart.SetEnabled(true)
					btnChooseDir.SetEnabled(true)
					return
				}
			}
		}()

	})
	return controlLayout
}

func updateStopInfo(InfoAboutScanningFiles *appStruct.CustomLabel)string{
	if strings.HasSuffix(InfoAboutScanningFiles.Text(),"..."){
		return "Подождите идет завершение процесса сканирования."
	}
	if strings.HasSuffix(InfoAboutScanningFiles.Text(),".."){
		return "Подождите идет завершение процесса сканирования..."
	}
	if strings.HasSuffix(InfoAboutScanningFiles.Text(),"."){
		return "Подождите идет завершение процесса сканирования.."
	}
	return "Подождите идет завершение процесса сканирования."
}

//сюда нужно передать appcomp
func optionalInfo(guiC *appStruct.GuiComponent)*widgets.QVBoxLayout{
	infoLayout    := widgets.NewQVBoxLayout()

	infoLayout.Layout().AddWidget(guiC.FileProgress)
	infoLayout.Layout().AddWidget(guiC.InfoAboutScanningFiles) //сколько просканировано
	infoLayout.Layout().AddWidget(guiC.ScanningTimeInfo) //сколько просканировано
	infoLayout.Layout().AddWidget(guiC.StartDirectoryForScan)

	return infoLayout

}
















