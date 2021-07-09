package extract

import (
	"github.com/aglyzov/charmap"
	"github.com/softlandia/cpd"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const maxReadFileSize = 50*1024*1024

func isBigfile(filepath string)bool{
	stat,_ := os.Stat(filepath)
	if stat.Size() > maxReadFileSize {
		return true
	}
	return false
}

func retrieveWords(wf map[string]int)[]string{
	var words []string
	for k,_ := range wf{
		words = append(words,k)
	}
	return words
}

func readFromSmallFile(f *os.File,wf *map[string]int,codePage string){
	byteText, _ := ioutil.ReadAll(f)
	words := retrieveWords(*wf)
	*wf = initWordFrequency(words)

	if codePage == "UTF-8" {

		*wf = getFreq(string(byteText),*wf)

	}

	if codePage == "CP1251" ||  codePage =="KOI8-R" {

		b2,_ := charmap.ANY_to_UTF8(byteText, "CP1251")

		wf2 := initWordFrequency(words)
		wf2 = getFreq(string(b2),wf2)

		b2,_ = charmap.ANY_to_UTF8(byteText, "KOI8-R")
		wf1 := initWordFrequency(words)
		wf1 = getFreq(string(b2),wf1)

		len1 := 0
		len2 := 0


		for _,v := range wf1{
			if v > 0 {
				len1++
			}
		}

		for _,v := range wf2{
			if v > 0 {
				len2++
			}
		}


		if len2>len1{
			*wf = wf2
		} else if len1 > len2 {
			*wf = wf1
		} else {
			if codePage == "CP1251" {
				*wf = wf2
			} else {
				*wf = wf1
			}
		}

	}

}

func readFromBigFile(f *os.File,wf *map[string]int,codePage string){
	var offset int64

	words := retrieveWords(*wf)
	*wf = initWordFrequency(words)

	if codePage == "UTF-8" {
		for {
			byteText := make([]byte, maxReadFileSize)
			n, err := f.ReadAt(byteText, offset)
			getFreq(string(byteText),*wf)
			offset += int64(n)
			if err == io.EOF {
				break
			}
		}
	}

	if codePage == "CP1251" ||  codePage =="KOI8-R" {
		byteText := make([]byte, maxReadFileSize)
		f.ReadAt(byteText, offset)

		b2,_ := charmap.ANY_to_UTF8(byteText, "CP1251")

		wf2 := initWordFrequency(words)
		wf2 = getFreq(string(b2),wf2)

		b2,_ = charmap.ANY_to_UTF8(byteText, "KOI8-R")
		wf1 := initWordFrequency(words)
		wf1 = getFreq(string(b2),wf1)

		len1 := 0
		len2 := 0


		for _,v := range wf1{
			if v > 0 {
				len1++
			}
		}

		for _,v := range wf2{
			if v > 0 {
				len2++
			}
		}


		if len2>len1{
			*wf = wf2
		} else if len1 > len2 {
			*wf = wf1
		} else {
			if codePage == "CP1251" {
				*wf = wf2
			} else {
				*wf = wf1
			}
		}

		offset += maxReadFileSize

		for {
			byteText := make([]byte, maxReadFileSize)
			n, err := f.ReadAt(byteText, offset)
			getFreq(string(byteText),*wf)
			offset += int64(n)
			if err == io.EOF {
				break
			}
		}
	}

}

func initWordFrequency(words []string)map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[strings.ToLower(word)] = 0
	}
	return wordFreq
}



func getFreq(startText string, wf map[string]int)map[string]int {

	startText = strings.ToLower(startText)
	words := retrieveWords(wf)
	for _, word := range words {
		word = strings.ToLower(word)
		var TextForSearch = startText
		end := true
		for end {
			index := strings.Index(TextForSearch, word)
			if index >= 0 {
				wf[word]++
				TextForSearch = TextForSearch[index+len(word):]
			} else {
				end = false
			}
		}
	}
	return wf
}

func GetTxtWordFrequency(filepath string,words []string) map[string]int{
	wf := initWordFrequency(words)
	codePage,_ := cpd.FileCodepageDetect(filepath)
	f, err := os.Open(filepath)

	if err != nil {
		return nil
	}
	defer f.Close()

	if isBigfile(f.Name()){
		readFromBigFile(f,&wf,codePage.String())
	} else {
		 readFromSmallFile(f,&wf,codePage.String())
	}


	return wf
}

func GetStringWordFrequency(text string,words []string) map[string]int{
	wf := initWordFrequency(words)
	wf = getFreq(text,wf)
	return wf
}



