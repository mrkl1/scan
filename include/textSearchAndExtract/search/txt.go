package search

import (
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract"
	"runtime/debug"
)

func recovery(st chan map[string]int) {
	if r := recover(); r != nil {
		debug.PrintStack()
		 st<-nil
	}
}

func Txt(path string,words []string,st chan map[string]int) {
	defer recovery(st)
	stat := extract.GetTxtWordFrequency(path,words)


	st<-stat

}