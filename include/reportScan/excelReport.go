package reportScan

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/therecipe/qt/widgets"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

var (
	filesSheetName = "Найденные файлы"
	errorsListName = "Список ошибок"
	mediaListName      = "Список медиа файлов"

	reducedCoefPath = 1.3
	increasedCoefWords = 1.1

	maxPathLength = 120.0
)

func SaveReport(guiC *appStruct.GuiComponent){

	var fileDialog = widgets.NewQFileDialog2(nil,"Save as...",
		"","")
	fileDialog.SetMimeTypeFilters([]string{"*.xlsx",})
	newPAth := fileDialog.GetSaveFileName(nil,"Save as...",
		"","Files (*.xlsx)","",0)



	f := excelize.NewFile()

	// Создаем три основных листа под отчет
	f.NewSheet(filesSheetName)
	f.NewSheet(errorsListName)
	f.NewSheet(mediaListName)
	//удаляется лист который создается при инициализации
	//нового файла
	f.DeleteSheet("Sheet1")

	//установка высоты по умолчанию
	//для первой строки 20
	//для последующих 15 или 15*n(где n кол-во строк)
	setFirstRowHeight(f,20)
	createFindFileList(f,guiC)
	createErrorList(f,guiC)
	createMediaList(f,guiC)

	f.SaveAs(newPAth)
}

func setFirstRowHeight(f *excelize.File,h float64){
	//можно красивее сделать но смысл
	//если тут врядли когда-нибудь будет больше 5 страниц
	f.SetRowHeight(filesSheetName,1,h)
	f.SetRowHeight(errorsListName,1,h)
	f.SetRowHeight(mediaListName,1,h)
}


func createFindFileList(f *excelize.File,guiC *appStruct.GuiComponent){

	f.SetCellValue(filesSheetName, "A1","Путь к файлам")
	f.SetCellValue(filesSheetName, "B1", "Частота вхождения слов")
	f.SetCellValue(filesSheetName, "C1", "Тип файла")

	f.SetColWidth(filesSheetName, "A","A", getRuneCount("Путь к файлам"))
	f.SetColWidth(filesSheetName, "B","B", getRuneCount("Частота вхождения слов")*increasedCoefWords)
	f.SetColWidth(filesSheetName, "C","C", getRuneCount("Тип файла")*increasedCoefWords)

	var lastInsertedID = 2

	topLvlCount := guiC.FileTree.TopLevelItemCount()
	path := guiC.StartDirectoryName
	for i := 0; i < topLvlCount;i++  {
		headItem := guiC.FileTree.TopLevelItem(i)
		if headItem.ChildCount() == 0 {

			saveToXLSXTable(f,filesSheetName ,filepath.Join(path,headItem.Text(0)),
				headItem.Text(1),headItem.Text(2),lastInsertedID)
				lastInsertedID++
		} else {
			path = filepath.Join(path,headItem.Text(0))
			findChildren(headItem,path,lastInsertedID,f)
		}
	}

}

func createErrorList(f *excelize.File,guiC *appStruct.GuiComponent){
	f.SetCellValue(errorsListName, "A1","Путь к файлам")
	f.SetCellValue(errorsListName, "B1", "Ошибки")
	f.SetCellValue(errorsListName, "C1", "Тип файла")

	f.SetColWidth(errorsListName, "A","A", getRuneCount("Путь к файлам"))
	f.SetColWidth(errorsListName, "B","B", getRuneCount("Ошибки")*increasedCoefWords)
	f.SetColWidth(errorsListName, "C","C", getRuneCount("Тип файла")*increasedCoefWords)

	var lastInsertedID = 2

	rowCount := guiC.ErrorTable.RowCount()
	for i := 0; i < rowCount;i++  {
			saveToXLSXTable(f,errorsListName ,guiC.ErrorTable.Item(i,0).Text(),
				guiC.ErrorTable.Item(i,1).Text(),guiC.ErrorTable.Item(i,2).Text(),lastInsertedID)
			lastInsertedID++
	}

}

func createMediaList(f *excelize.File,guiC *appStruct.GuiComponent){
	f.SetCellValue(mediaListName, "A1","Путь к файлам")
	f.SetCellValue(mediaListName, "B1", "Тип файла")

	f.SetColWidth(mediaListName, "A","A", getRuneCount("Путь к файлам"))
	f.SetColWidth(mediaListName, "B","B", getRuneCount("Тип файла")*increasedCoefWords)

	var lastInsertedID = 2

	rowCount := guiC.ErrorTable.RowCount()
	for i := 0; i < rowCount;i++  {
		saveToXLSXTable(f,mediaListName ,guiC.NonScanTable.Item(i,0).Text(),
			guiC.NonScanTable.Item(i,1).Text(),"",lastInsertedID)
		lastInsertedID++
	}

}

func findChildren(head *widgets.QTreeWidgetItem,path string,lastInsertedID int,f *excelize.File){
	for i := 0; i < head.ChildCount();i++  {
		newHead :=head.Child(i)
		if newHead.ChildCount() == 0 {
			saveToXLSXTable(f,filesSheetName ,filepath.Join(path,newHead.Text(0)),
				newHead.Text(1),newHead.Text(2),lastInsertedID)
			lastInsertedID++
		} else {
			path = filepath.Join(path,newHead.Text(0))
			findChildren(newHead,path,lastInsertedID,f)
		}
	}
}

func saveToXLSXTable(f *excelize.File,sheetName,path,words,ext string,lastInsertedID int){
		setCellParameters(lastInsertedID,sheetName,path,words,[]string{"A","B","C"},f)

		f.SetCellRichText(sheetName,"A"+fmt.Sprintf("%d",lastInsertedID),getRichText(path))
		f.SetCellRichText(sheetName,"B"+fmt.Sprintf("%d",lastInsertedID),getRichText(words))
		f.SetCellRichText(sheetName,"C"+fmt.Sprintf("%d",lastInsertedID),getRichText(ext))
}


func setCellParameters(rowIndex int,sheetName,filePath,words string,columns []string,f *excelize.File){
	curColWidth,_ := f.GetColWidth(sheetName,columns[0])

	if curColWidth <  getRuneCount(filePath)/reducedCoefPath {

		if getRuneCount(filePath)/reducedCoefPath > maxPathLength{
			f.SetColWidth(sheetName,columns[0],columns[0],maxPathLength)
		} else {
			f.SetColWidth(sheetName,columns[0],columns[0],getRuneCount(filePath)/reducedCoefPath)
		}

	}
	lineFeedCount := strings.Count(words,"\n")

	f.SetRowHeight(sheetName,rowIndex,(float64(lineFeedCount+1))*15)

	curBWidth,_ := f.GetColWidth(sheetName,columns[1])
	maxLen := 0.0

	if strings.Contains(words,"\n"){
		splitWords := strings.Split(words,"\n")

		for _,splWord := range splitWords {

			if float64(len(splWord)) > maxLen {
				maxLen = getRuneCount(splWord)

			}
		}
	} else {
		maxLen = getRuneCount(words)
	}

	if curBWidth < float64(maxLen) {

		f.SetColWidth(sheetName,columns[1],columns[1],maxLen)
	}
}

func getRuneCount(str string)float64{
	return float64(utf8.RuneCount([]byte(str)))
}

func getRichText(text string)[]excelize.RichTextRun {
	richTextRun := []excelize.RichTextRun{
		{
			Text: text,
			Font: &excelize.Font{
				Bold:      true,
				Italic:    false,
				Underline: "",
				Family:    "Times New Roman",
				Size:      10,
				Strike:    false,
				Color:     "2354e8",
			},
		},
	}
	return richTextRun
}

