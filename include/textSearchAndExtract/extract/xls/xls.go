package xls

import (
	"bytes"
	"fmt"
	"github.com/IntelligenceX/fileconversion"
	"log"
	"os"
)

func recoverAll() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}



func RetrieveTextFromXLS(filename string)string{
	//xlFile, err := xls.Open(filename, "utf-8")
	//if err != nil {
	//	return ""
	//}
	//defer xlFile.Close()
	//content := xlFile.ReadAllCells(10000)
	//
	//var text string
	//for _, f1 := range content {
	//	text += strings.Join(f1, " ")
	//}
//defer recoverAll()
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()
	stat, _ := file.Stat()
	buffer := bytes.NewBuffer(make([]byte, 10*1024*1024))

	_,err = fileconversion.XLS2Text(file,buffer,stat.Size())
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return ""
	}

	return buffer.String()
}
