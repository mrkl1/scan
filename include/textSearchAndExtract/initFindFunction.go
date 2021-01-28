package textSearchAndExtract



import (
	"github.com/myProj/scaner/new/include/textSearchAndExtract/search"
)

func FindText(path,ext string,words []string)map[string]int{
	finder,ok := m[ext]
	if ok{
		return finder(path,words)
	}
	return nil
}


type wordFrequencyFunc func(string,[]string)map[string]int

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
