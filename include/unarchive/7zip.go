package unarchive

import (
	"errors"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/saracen/go7z"
	"io"
	"log"
	"os"
	"path/filepath"
)

//TODO
// -сделать тесты где будут разные архивы находиться друг в друге

const (
	pass7zLinux = "7z"
	pass7zWindows = "7z/7z.exe"
)

func Unpack7z(path,dest string)error {

	if checkForPassword(path,".7z",nil){
		return errors.New("на файле пароль")
	}

	sz, err := go7z.OpenReader(path)
	if err != nil {
		logggerScan.SaveToLog("Unpack7z:"+err.Error())
		return errors.New("Не получилось открыть архив")
	}
	defer sz.Close()
	err = os.Mkdir(dest,0777)
	if err != nil {
		log.Println("Mkdir 7z error:",err)
		return errors.New("Не получилось открыть архив")
	}


	for {
		hdr, err := sz.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return errors.New("Не получилось открыть архив")
			//panic(err)
		}

		// If empty stream (no contents) and isn't specifically an empty file...
		// then it's a directory.
		if hdr.IsEmptyStream && !hdr.IsEmptyFile {
			if err := os.MkdirAll(filepath.Join(dest,hdr.Name) , os.ModePerm); err != nil {
				panic(err)
			}
			continue
		}

		// Create file
		f, err := os.Create(filepath.Join(dest,hdr.Name))
		if err != nil {
			log.Println(err)
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, sz); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}