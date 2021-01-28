package newWordsConfig

import (
	"github.com/myProj/scaner/new/include/logggerScan"
	"log"
	"path/filepath"
)

//сигнатура которая должна стоять в начле файле
var signature = "!scanner_word"
var dictionaryFilename = filepath.Join("config","dict.txt")
/*
Должен быть файл с сигнатурой
слова должны считываться построчно

нужно тримить строку перед запись в файл?

нужна возможность объединения файлов импорта экспорта

сделать защиту на случай если в строке будет слишком много слов (сам пользователь решает сколько
должно быть слов)

сделать настройку для того, чтобы игнорировать
пробелы в начале конце и ставить между словами по 1 пробелу

 */


func MergeImports(fileName string){
	if err := IsDictionaryFile(fileName); err!=nil {
		log.Println("MergeImports (IsDictionaryFile) error",err)
		return
	}
	words, err := ReadNewImport(fileName)
	if err != nil {
		logggerScan.SaveToLog("MergeImports (readFile) error:"+err.Error())
		log.Println("MergeImports (readFile) error",err)
		return
	}

	for _,w := range words {
		err := AddNewWord(w)
		if err != nil {
			logggerScan.SaveToLog("MergeImports add word error:"+err.Error())
		}
	}
}


