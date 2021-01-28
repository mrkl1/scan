package extract

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/myProj/scaner/new/include/logggerScan"
	"strings"
)

func RetriveTextFromXLSX(filename string)string{
	f, err := excelize.OpenFile(filename)
	if err != nil {
		logggerScan.SaveToLog("RetriveTextFromXLSX open file "+err.Error())
		fmt.Println(err)
	}

	sc := f.SheetCount
	//if sc == 0 {
	//	fmt.Println("sc = 0")
	//}
	// word = strings.ToUpper(word)
	var text string
	// var result bool
	for i := 0; i <= sc; i++ {

		textRow, _ := f.GetRows(f.GetSheetName(i))

		for _, f1 := range textRow {

			text += strings.Join(f1, " ")
		}
	}
	return text
}