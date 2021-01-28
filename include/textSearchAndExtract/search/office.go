package search

import (
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/docx"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/pptx"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/vsdx"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/extract/xls"
)

func Docx(path string,words []string)map[string]int{
	text,_ := docx.RetrieveAllText(path)
	return extract.GetStringWordFrequency(text,words)
}

func Xlsx(path string,words []string)map[string]int{
	text := extract.RetriveTextFromXLSX(path)
	return extract.GetStringWordFrequency(text,words)
}

func Vsdx(path string,words []string)map[string]int{
	vsdxDoc,_ := vsdx.ReadVSDXText(path)
	text := vsdxDoc.ExtractText()
	return extract.GetStringWordFrequency(text,words)
}

func Pptx(path string,words []string)map[string]int{
	pptxDoc,_ := pptx.ReadPPTXText(path)
	text := pptxDoc.ExtractText()
	return extract.GetStringWordFrequency(text,words)
}

func Xls(path string,words []string)map[string]int{
	text := xls.RetrieveTextFromXLS(path)
	return extract.GetStringWordFrequency(text,words)
}

func Doc(path string,words []string)map[string]int{
	text := extract.DocToTxt(path)
	return extract.GetStringWordFrequency(text,words)
}

func Rtf(path string,words []string)map[string]int{
	text := extract.Rtf2txt(path)
	return extract.GetStringWordFrequency(text,words)
}

func Pdf(path string,words []string)map[string]int{
	text := extract.ExtractTextFromPdf(path)
	return extract.GetStringWordFrequency(text,words)
}