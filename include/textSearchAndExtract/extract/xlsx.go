package extract

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/myProj/scaner/new/include/logggerScan"
	"log"
	"strings"
)

func RetriveTextFromXLSX(filename string)string{
	fmt.Println(filename)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		logggerScan.SaveToLog("RetriveTextFromXLSX open file "+err.Error())
		return ""
	}

	sc := f.SheetCount
	var text string
	// var result bool
	for i := 0; i <= sc; i++ {

		textRow, err := f.GetRows(f.GetSheetName(i))
		if err != nil {
			log.Println(err)
		}
		for _, f1 := range textRow {

			text += strings.Join(f1, " ")
		}
	}

	//PrintMemUsage()

	return text
}
















