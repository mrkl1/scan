package appStruct

import "github.com/therecipe/qt/widgets"


//структура хранит компоненты которые могут
//изменятся при работе программы
type GuiComponent struct{
	Application            *widgets.QApplication
	WordList 			   *widgets.QListWidget
	MainWindow             *widgets.QMainWindow
	MainWidget             *widgets.QWidget

	FileProgress           *CustomProgressBar
	FileProgressUpdate     chan int

	InfoAboutScanningFiles *CustomLabel


	StartDirectoryForScan  *CustomLabel

	ScanningTimeInfo       *CustomLabel

	FileTree               *CustomTreeWidget

	ErrorTable			   *widgets.QTableWidget


	NonScanTable		   *widgets.QTableWidget

	SearchIsActive         bool
	SkipItem               bool
	SkipItemNonArch               bool
	StartDirectoryName     string


	AddTempDir             chan string
	DeleteTempDir          chan string
	EndDeleteTemp          chan bool

	UpdateTime             string
	UpdateLabel            string
	ProgressBarValue       int


	EndUIUpdate              chan string
	ErrorTableUpdate         chan string
}

func NewGui()*GuiComponent{
	return &GuiComponent{
		Application:                  nil,
		WordList:                     nil,
		MainWindow:                   nil,
		MainWidget:                   nil,
		FileProgress:                 nil,
		FileProgressUpdate:           make(chan int, 1000),
		InfoAboutScanningFiles:       nil,
		StartDirectoryForScan:        nil,
		ScanningTimeInfo:             nil,
		FileTree:                     nil,
		ErrorTable:                   nil,
		ErrorTableUpdate:             make(chan string, 1000),
		NonScanTable:                 nil,
		EndUIUpdate:                  make(chan string, 2),
		SearchIsActive:               false,
		SkipItem:                     false,
		StartDirectoryName:           "",
		AddTempDir:                   make(chan string, 1000),
		DeleteTempDir:                make(chan string, 1000),
		EndDeleteTemp:                make(chan bool, 1000),
	}
}