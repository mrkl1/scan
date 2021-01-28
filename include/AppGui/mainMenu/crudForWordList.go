package mainMenu

import (
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/therecipe/qt/widgets"
)

func deleteFromWordList(wordList *widgets.QListWidget){
	itemsForRemove := wordList.SelectedItems()

	for _,item := range itemsForRemove {
		err := newWordsConfig.DeleteWord(item.Text())
		if err != nil {
			logggerScan.SaveToLog("deleteFromWordList error, word:"+item.Text()+"error:"+err.Error())
		}

	}
	updateWordList(wordList)

}

func addToWordList(word string,wordList *widgets.QListWidget)error{
	err := newWordsConfig.AddNewWord(word)
	updateWordList(wordList)
	return err
}

func updateWordList(wordList *widgets.QListWidget){
	wordList.Clear()
	wordList.AddItems(newWordsConfig.GetDictWords())
}