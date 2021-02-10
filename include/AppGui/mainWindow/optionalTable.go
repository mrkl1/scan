package mainWindow

import (
	"errors"
	"github.com/myProj/scaner/new/include/unarchive"
	"github.com/therecipe/qt/widgets"
)

func newErrorTable()*widgets.QTableWidget{
	table := widgets.NewQTableWidget(nil)
	//установить readOnly
	table.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	table.SetColumnCount(3)
	table.SetHorizontalHeaderLabels([]string{"Имя файла","Ошибки","Тип файла"})
	addErrorsToTable(table,unarchive.ArchInfoError{	ArchiveName: "",OpenError:   errors.New("test")},".test")
	table.RemoveRow(0)
	return table
}

func NewNonScanTable()*widgets.QTableWidget{
	table := widgets.NewQTableWidget(nil)
	//установить readOnly
	table.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	table.SetColumnCount(2)
	table.SetRowCount(0)
	table.SetHorizontalHeaderLabels([]string{"Имя файла","Тип файла"})
	addNonScanItem(table,".test",".test")
	table.RemoveRow(0)
	return table
}

func RemoveAllRows(table *widgets.QTableWidget){

	for table.RowCount()-1 > 0{

		table.RemoveRow(1)
	}
	if table.RowCount() > 0 {
		table.RemoveRow(table.RowCount()-1)
	}

}