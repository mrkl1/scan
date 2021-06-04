package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var settingsFileName = filepath.Join("config","settings.json")



type Settings struct {
	Setting        string `json:"setting"`
	IsAllowSetting bool   `json:"isAllowSetting"`
	Limit          int    `json:"limit"`
}



func ReadSettingsFromConfig()[]Settings{
	settingsFile,err := os.OpenFile(settingsFileName,os.O_RDWR | os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err,": settings config read")
		return nil
	}
	defer settingsFile.Close()
	settingsData,err := ioutil.ReadAll(settingsFile)

	var settings []Settings

	err = json.Unmarshal(settingsData,&settings)

	if err != nil {
		fmt.Println(err)
	}

	return settings
}

func SetNewConfig(s []Settings){
	newJson, err := json.MarshalIndent(&s, "", "  ")

	if err != nil {
		log.Println("Transformation error default file", err)
	}
	err = ioutil.WriteFile(settingsFileName, newJson, 0666)
	if err != nil {
		log.Println("Write file error default file:", err)
	}
}

//пока по сути просто заглушка для одной настройки
func IsNeedToShowError()bool{
	sets := ReadSettingsFromConfig()
	if len(sets) > 0 {
		if sets[0].Setting == "Отображать файлы в таблице с неизвестным расширением" {
			return sets[0].IsAllowSetting
		}
	}
	return true
}

func SetArchiveLimit(limit int){
	sets := ReadSettingsFromConfig()
	log.Println("sets",	sets)
	for i := 0;i  < len(sets); i++ {
		if sets[i].Setting == "Ограничение на минимальное свободное место при распаковке архивов в ГБ" {
			log.Println("sets[i]",	sets[i])
			sets[i].Limit = limit
		}
	}
	log.Println("sets",	sets)
	SetNewConfig(sets)
	return
}

func GetArchiveLimit()int{
	sets := ReadSettingsFromConfig()

	for _,s := range sets {
		if s.Setting == "Ограничение на минимальное свободное место при распаковке архивов в ГБ"{
			return s.Limit
		}
	}
	return 	1
}



