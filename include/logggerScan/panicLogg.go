package logggerScan

import (
	"log"
	"os"
)

const panicFileName = "panicStackTraceLog.log"

/*
PanicSave log to logfile
panic stacktrace
 */
func PanicSaveTrace(stackTrace string){
	f, err := os.OpenFile(panicFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening LOG file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(stackTrace)
	log.SetOutput(os.Stdout)
}
