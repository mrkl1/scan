package mainMenu

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const (
	//окно для добавления слов
	addWordWindowWidth  = 250
	addWordWindowMaxWidth  = 450
	addWordWindowHeight = 200
	addWordWindowName = "Добавление слов"
	addWordWindowInputName  =  "Введите слово"
	addWordWindowAddBtnName =  "&Добавить"
	//основное окно со списком слов
	wordListWindowWidth =  400
	wordListWindowHeight = 550
	wordListName       = "Слова для поиска"
	wordListWindowName = "Список слов"
	wordListWindowAddBtnName    = "&Добавить слово"
	wordListWindowDeleteBtnName = "&Удалить выделенные слова"
)

//newWordListWindow создает окно для отображения и редактирования
//списка слов
func newWordListWindow(guiC *appStruct.GuiComponent)*widgets.QWidget{
	//создание окна и начальная настройка
	wordListWindow := widgets.NewQWidget(nil,1)
	wordListWindow.SetWindowFlags(core.Qt__Dialog) //флаг поднимает окно над всеми
	wordListWindow.SetMinimumSize2(wordListWindowWidth,wordListWindowHeight)

	wordListWindow.SetWindowTitle(wordListWindowName)
	//для данного окна достаточно одного слоя (layout)
	vBoxLayout := widgets.NewQVBoxLayout()
	wordListWindow.SetLayout(vBoxLayout)
	//создание кнопок для редактирования списка слов
	btnAdd := widgets.NewQPushButton2(wordListWindowAddBtnName, nil)
	btnDelete := widgets.NewQPushButton2(wordListWindowDeleteBtnName, nil)
	//на данном виджете отображается список слов
	wordList := widgets.NewQListWidget(nil)
	//при таком флаге можно выделять несколько строк
	wordList.SetSelectionMode(2)
	wordList.SetWindowTitle(wordListName)
	wordList.AddItems(newWordsConfig.GetDictWords())
	//добавление компонентов на слой
	vBoxLayout.Layout().AddWidget(wordList)
	vBoxLayout.Layout().AddWidget(btnAdd)
	vBoxLayout.Layout().AddWidget(btnDelete)
	guiC.WordList = wordList
	addWordWindow := setAddWordWindow(wordList) // окно для добавления слов
	// вызов окна для добавления слов
	btnAdd.ConnectClicked(func(bool){
		addWordWindow.Show()
	})
	//удаляет слова из списка как на экране
	//так и из списка в конфиге
	btnDelete.ConnectClicked(func(bool){
		deleteFromWordList(wordList)
	})

	return wordListWindow
}

//setAddWordWindow вызвается из newWordListWindow
//создает окно для добавления списка слов
func setAddWordWindow(wordList *widgets.QListWidget)*widgets.QWidget{

	addWordWindow := widgets.NewQWidget(nil,0)
	vBoxLayout := widgets.NewQVBoxLayout()
	addWordWindow.SetLayout(vBoxLayout)

	addWordWindow.SetWindowFlags(core.Qt__Dialog) //флаг поднимает окно над всеми
	addWordWindow.SetMinimumSize2(addWordWindowWidth,addWordWindowHeight)
	addWordWindow.SetWindowTitle(addWordWindowName)
	addWordWindow.SetMaximumWidth(addWordWindowMaxWidth)
	//пользовательский ввод
	userInput := widgets.NewQLineEdit(nil)
	//кнопка для добавления слова в конфиг
	addBtn := widgets.NewQPushButton2(addWordWindowAddBtnName,nil)
	//область отображает было ли слово успешно добавлено в словарь или нет
	lblResult := widgets.NewQLabel2("", nil, 0)
	lblInputName := widgets.NewQLabel2(addWordWindowInputName,nil,0)
	lblInputName.SetFixedHeight(15)
	//добавление компонентов на форму
	vBoxLayout.Layout().AddWidget(lblInputName)
	vBoxLayout.Layout().AddWidget(userInput)
	vBoxLayout.Layout().AddWidget(addBtn)
	vBoxLayout.Layout().AddWidget(lblResult)
	//событие для добавления слова
	addBtn.ConnectClicked(func(bool){
		if userInput.Text() == ""{
			lblResult.SetText("Слово должно содержать\n хотя бы один символ")
			return
		}

		err := addToWordList(userInput.Text(),wordList)
		if err != nil {
			lblResult.SetText("Такое слово уже есть")
			return
		}
		lblResult.SetText("Слово "+userInput.Text()+".\nУспешно добавлено")
		userInput.Clear()

	})

	return addWordWindow
}