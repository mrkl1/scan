package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/therecipe/qt/widgets"
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
	btnChooseDir  := widgets.NewQPushButton2(btnChooseDirName,nil)
	btnStart      := widgets.NewQPushButton2(btnStartName,nil)
	btnStop       := widgets.NewQPushButton2(btnStopName,nil)
	btnSkip       := widgets.NewQPushButton2(btnSkipName,nil)
	controlLayout.Layout().AddWidget(btnChooseDir)
	controlLayout.Layout().AddWidget(btnSkip)
	controlLayout.Layout().AddWidget(btnStart)
	controlLayout.Layout().AddWidget(btnStop)
	btnStop.SetEnabled(false)
	btnStart.ConnectClicked(func(checked bool) {
		//возможно сюда нужно будет поместить
		//структуру для сохранения текущего прогресса
		guiC.SearchIsActive = true
		btnStart.SetEnabled(false)
		btnChooseDir.SetEnabled(false)
		btnStop.SetEnabled(true)
		//mb тут сделать select по остановке функции

		go renderFileTree(guiC,btnStart,btnChooseDir,btnStop)

	})

	btnSkip.ConnectClicked(func(checked bool){
		guiC.SkipItem = true
	})

	btnChooseDir.ConnectClicked(func(checked bool){
		filePath := widgets.QFileDialog_GetExistingDirectory(guiC.MainWindow,"Open Directory",guiC.StartDirectoryName,widgets.QFileDialog__DontUseNativeDialog)
		if filePath == "" {
			return
		}
		guiC.StartDirectoryForScan.SetText(directoryForScanning+": "+filePath)
		guiC.InfoAboutScanningFiles.SetText("")
		guiC.StartDirectoryName = filePath
	})

	btnStop.ConnectClicked(func(checked bool) {
		guiC.SearchIsActive = false
		btnStop.SetEnabled(false)
		go func (){
			guiC.InfoAboutScanningFiles.SetText("Подождите идет завершение процесса сканирования")
			for {
				time.Sleep(time.Millisecond*350)
				guiC.InfoAboutScanningFiles.SetText(updateStopInfo(guiC.InfoAboutScanningFiles))
				if guiC.SearchIsActive{
					guiC.InfoAboutScanningFiles.SetText("Процесс сканирования был остановлен")
					btnStart.SetEnabled(true)
					btnChooseDir.SetEnabled(true)
					return
				}
			}
		}()

	})
	return controlLayout
}

func updateStopInfo(InfoAboutScanningFiles *widgets.QLabel)string{
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
















