package appStruct

import "github.com/therecipe/qt/widgets"

/*
Возможно будет использоваться для того чтобы рендерить дерево
 */
type TreeItemsPair struct {
	Parent *widgets.QTreeWidgetItem
	Child *widgets.QTreeWidgetItem
}
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

	FileTree               *widgets.QTreeWidget
	FileTreeUpdate          chan TreeItemsPair

	ErrorTable			   *widgets.QTableWidget
	ErrorTableUpdate         chan string

	NonScanTable		   *widgets.QTableWidget
	NonScanTableUpdate     chan string

	SearchIsActive         bool
	SkipItem               bool
	StartDirectoryName     string


	AddTempDir             chan string
	DeleteTempDir          chan string
	EndDeleteTemp          chan bool
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
		FileTreeUpdate:               make(chan TreeItemsPair, 1000),
		ErrorTable:                   nil,
		ErrorTableUpdate:             make(chan string, 1000),
		NonScanTable:                 nil,
		NonScanTableUpdate:           make(chan string, 1000),
		SearchIsActive:               false,
		SkipItem:                     false,
		StartDirectoryName:           "",
		AddTempDir:                   make(chan string, 1000),
		DeleteTempDir:                make(chan string, 1000),
		EndDeleteTemp:                make(chan bool, 1000),
	}
}