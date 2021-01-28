/*
	В данном пакете содержатся функции
	по работе с  конфигурационным файлом
 */
package config

import "C"
import (
	"encoding/json"
	"errors"
	"github.com/myProj/scaner/new/include/config/subConfig"
	"io/ioutil"
	"os"
)


const (
	configFileName = "config/init.json"
)

//по хорошему нужна функция для установки текущего конфига
//и текущий конфиг уже будет возвращать нужная функция


type Config struct {
	IsCurrentConfig    bool
	ConfigName   	   string
	ColumnCount  	   int
	ColumnName   	   []string
	ColumnSize    	   []int
	RowHeight    	   int
	StartDir           string
	StartCoordinate    []int
}



func defaultConfig()(Config){
	return Config{
		IsCurrentConfig: true,
		ConfigName:      "Стандартный конфиг",
		ColumnCount:     0,
		ColumnName:      nil,
		ColumnSize:      nil,
		RowHeight:       0,
		StartDir:        "",
		StartCoordinate: []int{100,100,900,900},
	}
}

func GetCurrentConfig()Config{
	configs,err := ReadFromConfig()
	if err != nil {
		return defaultConfig()
	}
	for _,c := range configs{
		if c.IsCurrentConfig == true {
			c = setCorrectDataAt(c)
			return c
		}
	}
	return defaultConfig()
}


func setCorrectDataAt(cfg Config)Config{
	cfg.StartCoordinate = subConfig.CheckCorrectPosition(cfg.StartCoordinate)
	return cfg
}


//добавление новой записи в конфигурационный файл
func AddNewEntryToConfig(newConfig Config)error{

	configs,err := ReadFromConfig()
	if err != nil {
		return err
	}

	for _ ,cfg := range configs{
		if cfg.ConfigName == newConfig.ConfigName {
			return errors.New("имя конфига не должно повторяться")
		}
	}
	configs = append(configs,newConfig)

	err = SaveConfig(configs)
	return err
}

//удаление записи из конфигурационного файла
func RemoveEntryFromConfig(nameForRemove string)error{
	configs,err :=  ReadFromConfig()
	if err != nil{
		return err
	}
	var newConfig []Config
	for _,cfg := range configs{
		if cfg.ConfigName == nameForRemove{
			continue
		}
		newConfig = append(newConfig,cfg)
	}
	err = SaveConfig(newConfig)
	return err
}

//редактирование  конфига
func EditEntryInConfig(oldConfig,newConfig Config)error{
	err := RemoveEntryFromConfig(oldConfig.ConfigName)
	if err != nil {
		return err
	}
	err = AddNewEntryToConfig(newConfig)
	return err
}

//возвращает массив распакованныъ конфигов
//из конфигурацинного файла
func ReadFromConfig()([]Config,error){
	configFile,err := os.OpenFile(configFileName,os.O_RDWR | os.O_APPEND,0644)
	if err != nil {
		return nil,err
	}
	defer configFile.Close()

	configContent,err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil,err
	}

	var cfg []Config
	//при возникновении такой ошибки нужно будет выставлять стандартный конфиг
	err = json.Unmarshal(configContent,&cfg)
	return cfg,err
}

//производится запись конфига в файл
func SaveConfig(configs []Config)error{
	cfgBytes,err := json.MarshalIndent(configs,"","    ")
	ioutil.WriteFile(configFileName,cfgBytes,0666)
	return err
}

func SaveCurrentConfig(config Config)error{

	cfgs,_ := ReadFromConfig()

	for i,c := range cfgs{
		if c.IsCurrentConfig == true {
			cfgs[i]=config
		}
	}
	cfgBytes,err := json.MarshalIndent(cfgs,"","    ")
	err = ioutil.WriteFile(configFileName,cfgBytes,0666)
	return err
}