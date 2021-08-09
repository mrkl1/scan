package unarchive

import (
	"fmt"
	"github.com/aglyzov/charmap"
	"github.com/softlandia/cpd"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"
)

/*
WARNING
Возможно данную функцию можно будет усовершенствовать
при помощи перебора возможных русских кодировок,
например так, если после автоопределения кодировки
строка остается невалидной, то по очереди
перебираем возможные варианты, начиная с
chcp1251/866, KOI-8R и тд
 */
func RenameZipIncorrectName(zipPath string){
	fi,_ := ioutil.ReadDir(zipPath)
	for _,v := range fi {

		if !utf8.ValidString(v.Name()){
			fmt.Println("incorrect zip")
			cp := cpd.CodepageAutoDetect([]byte(v.Name()))
			text,_ := charmap.ANY_to_UTF8( []byte(v.Name()), cp.String())
			os.Rename(filepath.Join(zipPath,v.Name())  ,filepath.Join(zipPath,string(text)))
		}
	}
}

