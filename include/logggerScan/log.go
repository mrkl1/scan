package logggerScan

import (
	"log"
	"os"
	"strings"
)

func SaveToLog(logText string){
	f, err := os.OpenFile("config/logs.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening LOG file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(strings.Replace(logText,"\n","",-1))
	log.SetOutput(os.Stdout)
}

func RemoveLogs(){
	os.Remove("config/logs.log")
}