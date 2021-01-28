package unarchive

import (
	"log"
	"os"
	"strings"
)

/*
нужно запускать от админа т.к. нужно проставлять расширение
 */
func CheckExtension(path,ext string)string{

	if strings.HasSuffix(path,ext){
		return path
	}
	err := os.Rename(path, path+ext)
	if err != nil {
		log.Println(err)
	}
	return path+ext

}
