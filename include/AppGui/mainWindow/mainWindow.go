package mainWindow

import (
	"github.com/myProj/scaner/new/include/AppGui/mainMenu"
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)



//названия и надписи в главном окне
const(
	starSearch = "Начать поиск"
	directoryForScanning="Директория для сканирования"
	choseStartDirectory="Выбрать директорию для поиска"
	stopSearch ="Остановить поиск"
)

//размеры главного окна
const(
	minMainWindowHeight = 600
	minMainWindowWight = 600
)
/*
StartUI входная функция для начала работы приложения
отрисовываются все компоненты, определяются обработчики
событий
*/
func StartUI(){

	logggerScan.RemoveLogs()
	guiC := NewGuiComponent() //получение основных компонентов для формы
	setConfig(guiC) //установка начальных графических параметров

	mainLayout := widgets.NewQVBoxLayout()
	guiC.MainWidget.SetLayout(mainLayout)
	//данная область отвечает за отображение найденных файлов
	//в процессе сканирования
	mainLayout.AddLayout(newInfoAreaLayout(guiC),0)
	//в данной области располложены кнопки управления ходом сканирования
	mainLayout.AddLayout(newBottomArea(guiC),0)
	//установка главного меню
	guiC.MainWindow.SetMenuBar(mainMenu.NewBar(guiC))

	guiC.MainWindow.Show()

	//засунуть в функцию
	guiC.MainWindow.ConnectCloseEvent(func(event *gui.QCloseEvent) {
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




	guiC.Application.Exec()
}
//NewGuiComponent определяет компоненты которые будут
//находиться на главном окне
func NewGuiComponent()*appStruct.GuiComponent{
	guiC := appStruct.NewGui()
	guiC.Application  = widgets.NewQApplication(len(os.Args), os.Args)
	guiC.MainWindow   = widgets.NewQMainWindow(nil, 1)
	guiC.MainWidget   = widgets.NewQWidget(nil, 1)
	guiC.FileProgress = appStruct.NewCustomProgressBar(nil)
	guiC.FileTree     = newFileTree(guiC)
	guiC.ErrorTable   = newErrorTable()
	guiC.NonScanTable = NewNonScanTable()
	guiC.FileProgress.SetMinimum(0)
	guiC.FileProgress.SetMaximum(0)
	guiC.StartDirectoryForScan  = appStruct.NewCustomLabel2("",nil,0)
	guiC.InfoAboutScanningFiles  = appStruct.NewCustomLabel2("",nil,0)
	guiC.ScanningTimeInfo  = appStruct.NewCustomLabel2("",nil,0)
	//
	guiC.StartDirectoryForScan.SetFixedHeight(15)
	guiC.InfoAboutScanningFiles.SetFixedHeight(15)
	guiC.ScanningTimeInfo.SetFixedHeight(15)

	guiC.MainWindow.SetMinimumSize2(minMainWindowWight, minMainWindowHeight)
	guiC.MainWindow.SetCentralWidget(guiC.MainWidget)
	return guiC
}



