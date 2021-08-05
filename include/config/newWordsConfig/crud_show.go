package newWordsConfig

import (
	"bufio"
	"errors"
	"github.com/myProj/scaner/new/include/logggerScan"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//добавляет новое слово, если слово уже
//существует то сообщает об этом

func AddNewWord(word string)error{
	words,_ := ReadDictionary()
	if wordIsExist(word,words)>-1{
		return errors.New("word:"+word+" - already exist")
	}
	words = append(words,word)
	err := rewriteDictionary(words)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func DeleteWord(word string)error{
	words,_ := ReadDictionary()
	if i := wordIsExist(word,words);i > -1{
		words = append(words[:i],words[i+1:]...)


		err := rewriteDictionary(words)
		if err != nil {
			log.Println(err)
		}

		return nil
	}

	return errors.New("word:"+word+" - not found")
}

//возвращает список всех слов из словаря


//проверяет есть ли слово в словаре
func wordIsExist(word string,words []string)int{
	for i,w := range words{
		if w == word {
			return i
		}
	}

	return -1
}

func IsDictionaryFile(fileName string)error{
	file,err := os.Open(fileName)
	if err != nil {
		log.Println("Is Dictionary Error:",err)
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if  strings.HasPrefix(line,signature)   {
		return nil
	}
	return errors.New("not a dictionary file")
}

//Читает словарь построчно и возвращает массив
//возвращает все включая первую строку с сигнатурой
func ReadDictionary()([]string,error) {

	file, err := os.Open(dictionaryFilename)
	if err != nil {
		logggerScan.SaveToLog("readDictionary Error:"+ err.Error())
		log.Println("readDictionary Error:",err)
		return nil, err
	}
	defer file.Close()

	return readFile(file)
}

func ReadNewImport(fileName string)([]string,error) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Println("readDictionary Error:", err)
		return nil, err
	}
	defer file.Close()

	return readFile(file)
}

func GetDictWords()([]string) {

	file, err := os.Open(dictionaryFilename)
	if err != nil {
		log.Println("readDictionary Error:", err)
		return nil
	}
	defer file.Close()

	w,err := readFile(file)
	if err != nil {
		log.Println("GetDictWords",err)
	}
	return w[1:]
}


func readFile(file *os.File)([]string,error){
	reader := bufio.NewReader(file)
	var err error
	var line string
	var words []string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(line)>2{
			words = append(words, strings.TrimSpace(line))
		}
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		logggerScan.SaveToLog(" readFile > Failed with error:"+err.Error())
		log.Printf(" readFile > Failed with error: %v", err)
		return nil, err
	}

	return words, nil
}

//перезапись файла при добавлении/удалении слова
func rewriteDictionary(words []string)error{
	err := os.Remove(dictionaryFilename)
	if err != nil {
		logggerScan.SaveToLog(" rewriteDictionary error:"+err.Error())
		log.Printf("rewriteDictionary: %v", err)
	}
	file, err := os.OpenFile(dictionaryFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return errors.New("rewrite Error: "+err.Error())
	}
	defer file.Close()
	var dataS string
	for _, data := range words {
		data = strings.TrimSpace(data)
		if data != ""{
			dataS+=data+"\r\n"
		}

	}


	err = ioutil.WriteFile(dictionaryFilename,[]byte(dataS),0777)
	if err != nil {
		logggerScan.SaveToLog(" rewriteDictionary WriteFile error:"+err.Error())
		log.Println(err)
	}
	return nil
}



