package words

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var wordsFileName = filepath.Join("config","words.json")
//Структура для работы со словами
//в формате json
type Words struct {
	Word string `json:"word"`
}

func ReadWordsFromConfig()[]string{
	wordsFile,err := os.OpenFile(wordsFileName,os.O_RDWR | os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err,": words config read")
		return nil
	}
	defer wordsFile.Close()
	wordsData,_ := ioutil.ReadAll(wordsFile)

	var words []Words
	var listOfWords []string

	json.Unmarshal(wordsData,&words)


	for _,word := range words{
		listOfWords = append(listOfWords,word.Word)
	}

	return listOfWords
}

func AddWordToConfig(newWord string)error{
	if newWord == ""{
		return nil
	}
	wordsFile,err := os.OpenFile(wordsFileName,os.O_RDWR | os.O_APPEND,0644)

	if err != nil {
		fmt.Println(err,": words config add")
		return err
	}
	defer wordsFile.Close()
	wordsData,_ := ioutil.ReadAll(wordsFile)
	var words []Words
	err = json.Unmarshal(wordsData, &words)
	if err != nil {
		fmt.Println(err,": words json add")
		return err
	}
	for _, word := range words {
		if word.Word == newWord {
			fmt.Println(errors.New("word already exist"))
			return errors.New("word already exist")
		}
	}
	words = append(words,Words{Word: newWord})
	newJsonWords,_ := json.MarshalIndent(&words,""," ")
	ioutil.WriteFile(wordsFileName,newJsonWords,0666)
	return err
}

func DeleteWordFromConfig(newWord string)error{
	wordsFile,err := os.OpenFile(wordsFileName,os.O_RDWR | os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err,": words config delete")
		return err
	}
	defer wordsFile.Close()
	wordsData,_ := ioutil.ReadAll(wordsFile)
	var words []Words
	err = json.Unmarshal(wordsData, &words)
	if err != nil {
		fmt.Println(err,": words json delete")
		return err
	}
	for i, word := range words {
		if word.Word == newWord {
			words = append(words[:i],words[i+1:]...)

		}
	}
	newJsonWords,_ := json.MarshalIndent(&words,""," ")
	ioutil.WriteFile(wordsFileName,newJsonWords,0666)
	return err
}

func ImportJson(filename string){
	wordsFile,_ := os.OpenFile(filename,os.O_RDWR | os.O_APPEND,0644)
	defer wordsFile.Close()
	wordsData,_ := ioutil.ReadAll(wordsFile)
	var words []Words

	json.Unmarshal(wordsData,&words)

	for _,word := range words{
		AddWordToConfig(word.Word)
	}
}



