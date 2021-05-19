package unarchive

import (
	"bytes"
	"github.com/alexmullins/zip"
	"log"
	"os/exec"
	"strings"

)

func checkForPassword(path,ext string,err error)bool{
	switch ext {
	case ".zip":
		return checkZip(path,err)
	case ".rar":
		return checkRar(path,err)
	case ".7z":
		return check7z(path)
	}
	return false
}


func checkZip(path string,err error)bool{
	//открыть архив

	//проверка на пароль
	or, err := zip.OpenReader(path)
	if err != nil {
		return false
	}
	defer or.Close()
	if or.File[0].IsEncrypted() {
		return true
	}

	return false

}

func checkRar(path string,err error)bool{
	if err.Error() == "reading file in rar archive: rardecode: incorrect password" {
		return true
	}
	return false
}

//true password exist
func check7z(path string)bool{
	//
	var cmd *exec.Cmd

	cmd = getCommandPassword(path)

	//var out bytes.Buffer
	var stderr bytes.Buffer
	//cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(),"Can not open encrypted archive. Wrong password?") {
			log.Println(path,"Contains password")
			return true
		}
	}
	return false
}

//добавить 7z в path
func addToPATH(){
	//fmt.Println(os.Getenv("TEST_ENV_FOR_PROD"))
	//os.Setenv("TEST_ENV_FOR_PROD","/w/e/r/t:d/d/g/s")
	//old := os.Getenv("TEST_ENV_FOR_PROD")
	//fmt.Println(os.Getenv("TEST_ENV_FOR_PROD"))
	//newenv := old + ":" + "/w/g/a/c/g"
	//os.Setenv("TEST_ENV_FOR_PROD",	newenv)
	//fmt.Println(os.Getenv("TEST_ENV_FOR_PROD"))
}



