package search

import (
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/docx"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/pptx"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/vsdx"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/xls"
)


func Docx(path string,words []string,st chan map[string]int){
	text,_ := docx.RetrieveAllText(path)
	stat := extract.GetStringWordFrequency(text,words)
	st <- stat
}

func Xlsx(path string,words []string,st chan map[string]int){
	defer recoveryXLSX(path,words,st)
	text := extract.RetriveTextFromXLSX(path)
	stat :=  extract.GetStringWordFrequency(text,words)
	st <- stat

}




func Vsdx(path string,words []string,st chan map[string]int){
	vsdxDoc,_ := vsdx.ReadVSDXText(path)
	text := vsdxDoc.ExtractText()
	stat := extract.GetStringWordFrequency(text,words)
	st <- stat
}

func Pptx(path string,words []string,st chan map[string]int){
	pptxDoc,_ := pptx.ReadPPTXText(path)
	text := pptxDoc.ExtractText()
	stat := extract.GetStringWordFrequency(text,words)
	st <- stat
}

func Xls(path string,words []string,st chan map[string]int){
	text := xls.RetrieveTextFromXLS(path)
	stat := extract.GetStringWordFrequency(text,words)
	st <- stat
}

func Doc(path string,words []string,st chan map[string]int){
	text := extract.DocToTxt(path)
	stat :=  extract.GetStringWordFrequency(text,words)
	st <- stat
}

func Rtf(path string,words []string,st chan map[string]int){
	text := extract.Rtf2txt(path)
	stat :=  extract.GetStringWordFrequency(text,words)
	st <- stat
}

func Pdf(path string,words []string,st chan map[string]int){
	text := extract.ExtractTextFromPdf(path)
	stat := extract.GetStringWordFrequency(text,words)
	st <- stat
}