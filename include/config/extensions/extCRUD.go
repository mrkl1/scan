package extensions

import (
	"encoding/json"
	"errors"
	"github.com/myProj/scaner/new/include/logggerScan"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)


var (
	//ошибки
	openFileError =  errors.New("ошибка при открытии файла")
	//файл конфигурации
	extConfigFileName = filepath.Join("config","extConfig.json")
)

type Extension struct {
	Ext string `json:"ext"`
	AllowScanning bool `json:"allowScanning"`
}

type EXTensions []Extension


//Возвращает слайс разрешенных
//расширений для анализа
func GetAllowExtensions()([]string,error){
	allExtensions,err := readExtConfig()
	var allowedExtensions []string
	if err != nil {
		//TODO при ошибке возвращается автосгенеренный список
		log.Println("Ошибка при получении данных из конфига:", err)
		allExtensions = GetDefaultAllowList()
	}


	for _,e := range allExtensions{
		if e.AllowScanning{
			allowedExtensions = append(allowedExtensions,e.Ext)
		}
	}
	return allowedExtensions,err
}

//Считывает информацию из файла
//extConfig, преобразует данные в
//слайс Extension
func readExtConfig()([]Extension,error){
	var extensions []Extension
	//TODO тут должна быть проверка на существование файла
	// -если его нет то возвращается дефолтный слайс, где
	// разрешены все расширения, при этом создается новый
	// конфиг в котором все это будет разрешено
	// функция например checkFileExist

	extF,err := os.OpenFile(extConfigFileName,os.O_RDONLY,0777)
	if err != nil {
		log.Println("Возникла ошибка при открытии файла",extConfigFileName,err)
		logggerScan.SaveToLog("Возникла ошибка при открытии файла "+extConfigFileName+err.Error())
		return nil,openFileError
	}
	defer extF.Close()

	data, err := ioutil.ReadAll(extF)
	if err != nil {
		log.Println("Возникла ошибка при чтении данных из файла",extConfigFileName,err)
		logggerScan.SaveToLog("Возникла ошибка при чтении данных из файла "+extConfigFileName+" "+err.Error())
		return nil,openFileError
	}

	err = json.Unmarshal(data,&extensions)
	if err != nil {
		log.Println("Ошибка при маршалинге данных из файла конфигурации "+extF.Name()+" "+err.Error())
		return nil,openFileError
	}

	return extensions,nil

}

//позволяет запретить/разрешить сканировать
//файлы с выбранным расширением
//todo
//	возможно тут нужно будет сделать обертку над этой функцией
// 	чтобы в цикле проставляло нужные значения для всех аргументов (расширений)
func SetExtStatus(extension []Extension)error{

	err := writeNewConfig(extension)
	if err != nil {
		logggerScan.SaveToLog("SetExtStatus error:"+err.Error())
	}
	return err
}