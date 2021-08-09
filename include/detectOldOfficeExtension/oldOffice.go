package detectOldOfficeExtension

import (
	"bytes"
	"github.com/richardlehane/mscfb"
	"io"
	"os"
)

func isOle(in []byte) bool {
	return bytes.HasPrefix(in, []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1})
}

/*
Работает гораздо лучше чем от vasilie mimetype, но все равно не идеально
2021/08/05 16:20:10 /home/us/Загрузки/7z/test.doc[ext:.ole]
 */
func DetectOldOffice(filename string) string {

	file, err := os.Open(filename)

	if err != nil {

		return ""
	}

	sign := make([]byte,8)
	//file.Read зависает и не дает идти дальше
	_,err = file.ReadAt(sign,0)

	if err != nil {

		return ""
	}

	if !isOle(sign) {
		return ""
	}
	_,err = file.Seek(0,io.SeekStart)
	if err != nil {
		return ""
	}
	doc, err := mscfb.New(file)

	if err!= nil {
		return ""
	}

	    entry := getLastSummaryInformation(doc)

	    if entry == nil {
	    	return ""
		}

		if entry.Name == "SummaryInformation"{

			buf := make([]byte,entry.Size)
			_,err = entry.Read(buf)
			if err!= nil {

				return ""
			}
			if bytes.Contains(buf,[]byte("Microsoft Office PowerPoint")) || bytes.Contains(buf,[]byte("Microsoft PowerPoint")){
				return ".ppt"
			}
			if bytes.Contains(buf,[]byte("Microsoft Excel")) || bytes.Contains(buf,[]byte("Microsoft Office Excel")) {
				return ".xls"
			}
			if bytes.Contains(buf,[]byte("Microsoft Office Word")) || bytes.Contains(buf,[]byte("Microsoft Word")) {
				return ".doc"
			}
			if bytes.Contains(buf,[]byte("Microsoft Visio")) || bytes.Contains(buf,[]byte("Microsoft Office Visio")) {
				return ".vsd"
			}
		}



	return ".ole"
}

//особенность в том, что в документ могут быть встроены
//другие документы и из-за этого расширение мб неправильно
//распознано, поэтому берется последний файл
//SummaryInformation, который относится к основному документу (но это не точно)
//в целом пока не было документов которые бы ошибочно определялись этим способом
func getLastSummaryInformation(doc *mscfb.Reader)*mscfb.File{
	var lastEntry *mscfb.File
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {
		name := entry.Name
		if name == "SummaryInformation" {
			lastEntry = entry
		}
	}

	return lastEntry
}