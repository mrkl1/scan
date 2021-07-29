package extract

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/myProj/scaner/new/include/logggerScan"
	"log"
	"strings"
)

func RetriveTextFromXLSX(filename string)string{
	f, err := excelize.OpenFile(filename)
	if err != nil {
		logggerScan.SaveToLog("RetriveTextFromXLSX open file "+err.Error())
		return ""
	}

	sc := f.SheetCount
	var text string

	for i := 0; i <= sc; i++ {

		textRow, err := f.GetRows(f.GetSheetName(i))
		if err != nil && !strings.HasSuffix(err.Error(),"is not exist"){
			log.Println(err)
		}
		for _, f1 := range textRow {

			text += strings.Join(f1, " ")
		}
	}
	return text
}
















