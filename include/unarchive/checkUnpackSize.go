package unarchive

import (
	"bytes"
	"errors"
	mindisk "github.com/minio/minio/pkg/disk"
	"github.com/myProj/scaner/new/include/logggerScan"
	"log"
	"runtime"
	"strconv"
	"strings"
)

//limit for the file unpack
var unpackLimit uint64  = 500*1024*1024

func getFreeSpace(path string)(uint64,error){
	di, err := mindisk.GetInfo(path)
	return di.Free,err
}

func getUnpackedSize(path string)(uint64,error){
	cmd := getCommandCheckSize(path)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	//в 7z в 3 столбце хранится общая инфа о размере распакованного файла

	size := strings.Fields(readLastLine(out.String()))[2]
	s,_ := strconv.ParseUint(size,10,64)

	if err != nil {
		log.Println("checkSize error :::"+stderr.String(),err)
		return 0,errors.New("Не получилось определить размер файла")
	}
	return s,nil
}

func readLastLine(s string)string{
	li1 := strings.LastIndex(s,"\n")
	//win notepad
	if s[li1-1:li1] == "\r"{
		li2 := strings.LastIndex(s[:li1],"\n")
		return s[li2+1:li1-1]
	}
	//lin gedit
	li2 := strings.LastIndex(s[:li1],"\n")
	return s[li2+1:li1]
}

func IsSpaceEnough (archPath string)bool{
	size,_ := getUnpackedSize(archPath)

	partName,_ := partitionName(archPath)
	freeSpace,_ := getFreeSpace(partName)
	if  freeSpace - size < unpackLimit {
		return false
	} else {
		return true
	}
}

func partitionName(path string)(string,error){
	if runtime.GOOS == "linux" {
		return "/",nil
	}

	spl := strings.Split(path,"/")[0]
	if len(spl) < 1 {
		logggerScan.SaveToLog("Get part name error with path: "+path)
		return "",errors.New("Get part name error")
	}
	return spl,nil
}