package extract
//
//import (
//	"golang.org/x/text/encoding/charmap"
//	"io"
//	"io/ioutil"
//	"os"
//	"strings"
//)
//
//const maxReadFileSize = 50*1024*1024
////TODO
//// добавить кодировку для chcp 861
//// спросить про то адекватный ли это код
////расположены по частоте встречаемости
//var alphabet = []string{
//	"о","а","е","и","н","т","р","с","л","в",
//	"к","п","м","у","д","я","ы","ь","з","б","г",
//	"й","ч","ю","х","ж","ш","ц","щ","ф","э","ъ"}
////для проверки на utf8
//var incorrectAlphabet = []string{
//	"Рѕ","Р°","Рµ","Рё","РЅ","С‚","СЂ","СЃ","Р»","РІ","Рє","Рї","Рј",
//	"Сѓ","Рґ","СЏ","С‹","СЊ","Р·","Р±","Рі","Р№","С‡","СЋ","С…","Р¶",
//	"С€","С†","С‰","С„","СЌ","СЉ",
//}
//
//
//func isBigfile(filepath string)bool{
//	stat,_ := os.Stat(filepath)
//	if stat.Size() > maxReadFileSize {
//		return true
//	}
//	return false
//}
//
//func retrieveWords(wf map[string]int)[]string{
//	var words []string
//	for k,_ := range wf{
//		words = append(words,k)
//	}
//	return words
//}
//
//func readFromSmallFile(f *os.File,wf *map[string]int){
//	byteText, _ := ioutil.ReadAll(f)
//	check1251text(byteText,*wf)
//
//	if containRus(wf){
//
//		return
//	} else {
//
//		words := retrieveWords(*wf)
//		*wf = initWordFrequency(words)
//		getFreq(string(byteText),*wf)
//	}
//}
//
//func readFromBigFile(f *os.File,wf *map[string]int){
//	var offset int64
//
//	for {
//		byteText := make([]byte, maxReadFileSize)
//		n, err := f.ReadAt(byteText, offset)
//		check1251text(byteText,*wf)
//		offset += int64(n)
//
//		if err == io.EOF {
//			break
//		}
//	}
//
//	if containRus(wf) {
//		return
//	} else {
//
//		words := retrieveWords(*wf)
//		*wf = initWordFrequency(words)
//		offset = 0
//		for {
//			byteText := make([]byte, maxReadFileSize)
//			n, err := f.ReadAt(byteText, offset)
//			getFreq(string(byteText),*wf)
//			offset += int64(n)
//			if err == io.EOF {
//				break
//			}
//		}
//
//	}
//
//
//}
//
//func isRusWord(word string)bool{
//	for _,a := range alphabet {
//		if strings.Contains(word,a) {
//			return true
//		}
//	}
//	return false
//}
//
//func isBadRusText(text string)bool{
//	for _,a := range incorrectAlphabet {
//		if strings.Contains(text,a) {
//			return true
//		}
//	}
//	return false
//}
//func containRus(wc *map[string]int)bool{
//	for k,v := range *wc{
//		if isRusWord(k) && v > 0{
//			return true
//		}
//	}
//	return false
//}
//
//func initWordFrequency(words []string)map[string]int {
//	wordFreq := make(map[string]int)
//	for _, word := range words {
//		wordFreq[strings.ToLower(word)] = 0
//	}
//	return wordFreq
//}
//
//
//func check1251text(textUTF8 []byte,wf map[string]int){
//	dec := charmap.Windows1251.NewDecoder()
//	text1251, _ := dec.Bytes(textUTF8)
//	if isBadRusText(string(text1251)){
//		return
//	}
//
//	getFreq(string(text1251),wf)
//
//
//}
//
//func getFreq(startText string, wf map[string]int)map[string]int {
//
//	startText = strings.ToLower(startText)
//	words := retrieveWords(wf)
//	for _, word := range words {
//		word = strings.ToLower(word)
//		var TextForSearch = startText
//		end := true
//		for end {
//			index := strings.Index(TextForSearch, word)
//			if index >= 0 {
//				wf[word]++
//				TextForSearch = TextForSearch[index+len(word):]
//			} else {
//				end = false
//			}
//		}
//	}
//	return wf
//}
//
//func GetTxtWordFrequency(filepath string,words []string) map[string]int{
//	wf := initWordFrequency(words)
//	f, err := os.Open(filepath)
//	if err != nil {
//		return nil
//	}
//	defer f.Close()
//	if isBigfile(f.Name()){
//		readFromBigFile(f,&wf)
//		return nil
//	} else {
//		readFromSmallFile(f,&wf)
//	}
//
//
//	return wf
//}
//
//func GetStringWordFrequency(text string,words []string) map[string]int{
//	wf := initWordFrequency(words)
//	wf = getFreq(text,wf)
//	return wf
//}
//
//
//
