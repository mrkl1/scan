package pptx

import (
	"archive/zip"
	"io/ioutil"
	"strings"
)

const tagTextOpen  = "<a:t>"
const tagTextClose = "</a:t>"
const openAngle  = "<"
const closeAngle = "/>"

//Стктура для работы с документом visio
type DocPPTX struct {
	zipReader  *zip.ReadCloser
	Files      []*zip.File
	pagesContent [][]byte
}

//Закрыть документ
func (d *DocPPTX)Close()error{
	return d.zipReader.Close()
}

func ReadPPTXText(path string)(DocPPTX,error){
	reader,err := zip.OpenReader(path)
	if err!=nil{
		return DocPPTX{}, err
	}
	defer reader.Close()
	doc := DocPPTX{
		zipReader:  reader,
		Files:      reader.File,
		pagesContent: nil,
	}


	for _, f := range reader.File {
		if  strings.Contains(f.Name,"ppt/slides")  {
			doc.pagesContent = append(doc.pagesContent, getPageContent(f))
		}
	}
	return doc, nil

}

func getPageContent(f *zip.File)[]byte{
	file, _ := f.Open()
	c,_ := ioutil.ReadAll(file)
	return c
}

func (d *DocPPTX)ExtractText()string{

	var result string

	for _,page := range d.pagesContent {
		startIndex := 0
		text := string(page)
		for {
			text = text[startIndex:]
			indexOpen,indexClose := findIndexesOfTag(text, tagTextOpen, tagTextClose)
			if indexOpen == -1 || indexClose == -1 {
				break
			}

			curTag := text[indexOpen : indexOpen+len(tagTextOpen)]
			if isOpenTag(curTag){
				tagText := text[indexOpen+len(tagTextOpen):indexClose]
				result  += removeInsideTags(tagText)
				startIndex = indexClose + len(tagTextClose)
				continue
			}
		}

	}
	return result
}

func isOpenTag(tag string)bool{
	return tag == tagTextOpen
}

func removeInsideTags(text string)string{
	for {
		indexOpen,indexClose := findIndexesOfTag(text, openAngle, closeAngle)
		if indexOpen == -1 || indexClose == -1 {
			break
		}

		text = strings.Replace(text,text[indexOpen:indexClose+len(closeAngle)],"",1)
	}
	return text
}

func findIndexesOfTag(text,openTag,closeTag string)(indexOpen,closeIndex int){
	return strings.Index(text, openTag),strings.Index(text, closeTag)
}



