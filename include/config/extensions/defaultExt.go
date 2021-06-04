package extensions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)



//Генерит список который разрешает сканирование всех
//типов файлов
func GetDefaultAllowList()[]Extension{
	var exts = make([]Extension,len(IncludeExtension))
	for i := 0; i < len(IncludeExtension);i++{
		exts[i].Ext = IncludeExtension[i]
		exts[i].AllowScanning = true
	}
	sort.Sort(EXTensions(exts))
	return  exts
}

func checkCorrectConfig(configExt []Extension)[]Extension{
	dExtList := GetDefaultAllowList()

	if len(configExt) == len(dExtList){
		return configExt
	}

	var newExtMap = make(map[string]bool,len(dExtList))
	for _,e := range dExtList{
		newExtMap[e.Ext] = true
	}

	for _,e := range configExt {
		newExtMap[e.Ext] = e.AllowScanning
	}
	configExt = nil
	for k,v := range newExtMap {
		configExt = append(configExt,Extension{
			Ext:           k,
			AllowScanning: v,
		})
	}
	return configExt
}

func GetAllowList()[]Extension{
	rec,_ := readExtConfig()
	rec = checkCorrectConfig(rec)

	sort.Sort(EXTensions(rec))

	return rec
}



/*
TODO такие ситуации помечать как critical и записывать в отдельный лог
 сделать свою структуру под это где будет сообщаться ошибка при этом
 если такая ошибка уже была то просто делать сокр запись об этом
 */
func CreateDefaultAllowListFile()error {
	var exts = make([]Extension, len(IncludeExtension))

	for i := 0; i < len(IncludeExtension); i++ {
		exts[i].Ext = IncludeExtension[i]
		exts[i].AllowScanning = true
	}

	return writeNewConfig(exts)
}

func writeNewConfig(e []Extension)error{
	newJson, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Println("Transformation error default file", err)
		return err
	}
	err = ioutil.WriteFile(extConfigFileName, newJson, 0666)
	if err != nil {
		log.Println("Write file error default file:", err)
		return err
	}
	return err
}




