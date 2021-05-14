package unarchive

import (
	"bytes"
	"errors"
	"fmt"
	mindisk "github.com/minio/minio/pkg/disk"
	"log"
	"strconv"
	"strings"
)

//limit for the file unpack
var unpackLimit uint64  = 500*1024*1024*1024

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
	//rar size
	//https://github.com/gen2brain/go-unarr
	size,_ := getUnpackedSize(archPath)


	freeSpace,_ := getFreeSpace(archPath)
	fmt.Println(freeSpace - size,unpackLimit)
	if  (freeSpace - size) < unpackLimit {
		println("not enough space")
		return false
	}

	return true
}
