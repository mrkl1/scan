package search

import (
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract"
)

func Txt(path string,words []string,st chan map[string]int) {
	stat := extract.GetTxtWordFrequency(path,words)

	st<-stat
}