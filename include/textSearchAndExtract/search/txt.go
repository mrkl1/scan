package search

import "github.com/myProj/scaner/new/include/textSearchAndExtract/extract"

func Txt(path string,words []string)  map[string]int {
	return extract.GetTxtWordFrequency(path,words)
}