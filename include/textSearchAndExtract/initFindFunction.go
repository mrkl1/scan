package textSearchAndExtract



import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/textSearchAndExtract/search"
	"time"
)



func skip(guiC *appStruct.GuiComponent,skip chan bool,end chan bool){

	for {
		time.Sleep(time.Millisecond*50)
		if guiC.SkipItem == true || guiC.SearchIsActive == false{
			guiC.SkipItem = false
			skip <- false
			return
		}
		select {
		case <-end:
			return
		default:
			continue
		}

	}
}


func FindText(path,ext string,words []string,guiC *appStruct.GuiComponent)map[string]int{
	/*
	объявляем контекст во внешней функции
	затем отменяем
	 */

	skipC       := make(chan bool,5)
	end       := make(chan bool,5)

	go skip(guiC,skipC,end)

	finder,ok := m[ext]

	if ok{

		stat := make(chan map[string]int,0)
		go  finder(path,words,stat)

		select {
		case st := <- stat:
			end <- true
			return st
		case <-skipC:
			return make( map[string]int,0)
		}
	}
	return make( map[string]int,0)
}


type wordFrequencyFunc func(string,[]string,chan map[string]int)

var m map[string]wordFrequencyFunc
//для более сложных случаев
//type WordFrequency struct {
//	frequensiers map[string] wordFrequencyFunc
//}
//TODO csv,odt,ott,ods parser
// + еще форматы для опен офиса
// собрать TODO и посмотреть что нужно сделать
func init(){
	m = make(map[string]wordFrequencyFunc)
	//text files
	m[".txt"] = search.Txt
		//configs and markdown languages
		m[".json"] = search.Txt
		m[".xml"] = search.Txt
		m[".html"] = search.Txt
		//programming languages
		m[".py"] = search.Txt
		m[".js"] = search.Txt
		m[".php"] = search.Txt
		m[".pl"] = search.Txt
		m[".lua"] = search.Txt
		m[".tcl"] = search.Txt




	//office old
	m[".xls"] = search.Xls
	m[".doc"]  = search.Doc
	//office new
	m[".docx"] = search.Docx
	m[".xlsx"] = search.Xlsx
	m[".pptx"] = search.Pptx
	m[".vsdx"] = search.Vsdx
	//pdf rtf and other
	m[".pdf"]  = search.Pdf
	m[".rtf"]  = search.Rtf



}
