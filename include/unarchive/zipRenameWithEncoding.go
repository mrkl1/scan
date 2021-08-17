package unarchive

import (
	"archive/zip"
	"fmt"
	"github.com/aglyzov/charmap"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/softlandia/cpd"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
WARNING
Эта функция работает нормально только под линухом
в детали не вникал, но скорее всего связанно
с тем что винда мб неправильно интерпретирует символы
определенной кодировки.
 */
func RenameZipIncorrectName(zipPath string){
	fi,err := ioutil.ReadDir(zipPath)
	if err != nil {
		logggerScan.SaveToLog("RenameZipIncorrectName err:"+err.Error())
	}
	logggerScan.SaveToLog("******"+zipPath)
	for _,v := range fi {

		logggerScan.SaveToLog(zipPath+"+"+v.Name()+"+"+strconv.FormatBool(utf8.ValidString(v.Name()))+"+"+cpd.CodepageAutoDetect([]byte(v.Name())).String())
		if !utf8.ValidString(v.Name()){
			logggerScan.SaveToLog("incorrect zip")
			cp := cpd.CodepageAutoDetect([]byte(v.Name()))
			text,_ := charmap.ANY_to_UTF8( []byte(v.Name()), cp.String())
			logggerScan.SaveToLog("****** "+string(text)+" (name)")
			os.Rename(filepath.Join(zipPath,v.Name())  ,filepath.Join(zipPath,string(text)))
			os.Rename(filepath.Join(zipPath,v.Name())  ,filepath.Join(zipPath,string(text)))
		}
	}
}

//UnzipUTF8 анзипает в правильной кодировке, обычно
func UnzipUTF8(src string, dest string) error {


	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()



	for _, f := range r.File {

		fWriteName := f.Name
		if !utf8.ValidString(f.Name){

			cp := cpd.CodepageAutoDetect([]byte(f.Name))
			utf8Name,_ := charmap.ANY_to_UTF8( []byte(f.Name), cp.String())
			logggerScan.SaveToLog("incorrect zip "+string(utf8Name)+"(name)"+"enc"+cp.String())
			fWriteName = string(utf8Name)
		}

		fpath := filepath.Join(dest, fWriteName)


		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return  err
		}
	}
	return nil
}
