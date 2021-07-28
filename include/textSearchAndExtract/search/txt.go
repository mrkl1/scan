package search

import (
	"fmt"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract"
	"github.com/tealeg/xlsx"
	"math"
	"runtime/debug"
	"strings"
	"reflect"
)

func recovery(st chan map[string]int,filename string) {
	if r := recover(); r != nil {
		fmt.Println(string(debug.Stack()))

		var panicText = "Panic text:\n"
		if reflect.TypeOf(r).String() == "string"{
			panicText = fmt.Sprintf("%v\n", r)
		}

		logggerScan.PanicSaveTrace(filename+"\n"+panicText+"Stack TRACE:\n"+string(debug.Stack()))
		//нужен какой то флаг
		stat :=  make(map[string]int,1)
		stat["!PANIC!01"]= math.MaxInt32
		fmt.Println(stat)
		st<-stat
	}
}

func recoveryXLSX(path string,words []string,st chan map[string]int) {

	if r := recover(); r != nil {

		fmt.Println(string(debug.Stack()))

		var panicText = "Panic text:\n"
		if reflect.TypeOf(r).String() == "string"{
			panicText = fmt.Sprintf("%v\n", r)
		}

		logggerScan.PanicSaveTrace(path+"\n"+panicText+"Stack TRACE:\n"+string(debug.Stack()))
		tealegXLXS(path,words,st)

	}
}

func tealegXLXS(filename string,words []string,st chan map[string]int){
	wb,err := xlsx.OpenFile(filename)
	defer recovery(st,filename)
	if err != nil {
		fmt.Println(err,wb)
		st <- nil
		return
	}
	panic("TEST")
	var text strings.Builder
	for _, w := range wb.Sheets {

		for i := 0; i < w.MaxRow; i++  {
			for j := 0; j < w.MaxCol ; j++ {
				c,_ := w.Cell(i,j)
				text.WriteString(c.Value)
				text.WriteString(" ")

			}
		}
	}

	stat :=  extract.GetStringWordFrequency(text.String(),words)
	st <- stat
}

func Txt(path string,words []string,st chan map[string]int) {
	defer recovery(st,path)
	stat := extract.GetTxtWordFrequency(path,words)

	st<-stat

}