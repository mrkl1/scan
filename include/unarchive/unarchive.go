package unarchive

import (
	"context"
	"errors"
	"github.com/gabriel-vasile/mimetype"
	"github.com/mholt/archiver"
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/myProj/scaner/new/include/tempDeleter"
	"github.com/myProj/scaner/new/include/textSearchAndExtract"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)


type chanCLosers struct {
	cancel chan bool
	skip   chan bool
}

func newChanCLosers()chanCLosers{
	return chanCLosers{
		cancel: make(chan bool,1),
		skip:   make(chan bool,1),
	}
}




type FrequencyInfo struct {
	Name          string
	WordFrequency map[string]int
	Ext           string
}


type ArchInfoError struct {
	ArchiveName string
	OpenError   error
	Ext 	    string
}



//**TODO функция для подстановки расширений
var exts = []string{".tar",".tar.gz",".zip",".rar",".7z",".gz",".tar.xz"}

func isArch(ext string)bool{
	for _,e := range exts {
		if e == ext {
			return true
		}
	}

	return false
}


func addInfoOrError(freqInf *[]FrequencyInfo,freqErr *[]ArchInfoError,
	fi chan FrequencyInfo,fe chan ArchInfoError,endInfo chan bool){

	for {

		select {
		case newFi := <-fi:
			*freqInf = append(*freqInf,newFi)
		case newFe := <-fe:
			*freqErr = append(*freqErr,newFe)
		case <- endInfo:
			return
		}
	}
}

func UnpackWithCtx(path,ext,beautyName string,guiC *appStruct.GuiComponent)([]FrequencyInfo,[]ArchInfoError){
	ctx,_ := context.WithCancel(context.Background())

	var freqInf []FrequencyInfo
	var freqErr []ArchInfoError

	path       = CheckExtension(path,ext)
	fichan     := make(chan FrequencyInfo,100)
	fechan     := make(chan ArchInfoError,100)

	end        := make(chan bool,5)
	endInfo    := make(chan bool,5)
	cancel     := make(chan bool,5)
	skip       := make(chan bool,5)
	stopTimer  := make(chan bool,5)


	ncc := newChanCLosers()

	if !guiC.IsTimeUpdate{

		var mu sync.Mutex
		mu.Lock()
		guiC.IsTimeUpdate = true
		mu.Unlock()
		go setTimeEverySecond(guiC,stopTimer)
	}

	go cancelUnpack(guiC,cancel,ncc.cancel)
	go skipUnpack(guiC,skip,ncc.skip)
	go addInfoOrError(&freqInf,&freqErr,fichan,fechan,endInfo)
	go wrapperUnpackCtx(path,ext,beautyName,guiC,fichan,fechan,end,ctx)



	defer func() {

		ncc.cancel <- false
		ncc.skip <- false
		end <- false
		endInfo <- false
		cancel <- false
		skip <- false
		stopTimer <- false
		if globCMD != nil {
			if globCMD.Process != nil {
				globCMD.Process.Release()
				//globCMD.Process.Kill()
			}
		}
		//cancelCtx()
	}()
	select {
	case <-cancel:
		return freqInf,freqErr
	case <-skip:
		return freqInf,freqErr
	case <- end :
		time.Sleep(1*time.Millisecond)
		return freqInf,freqErr
	}

}

func cancelUnpack(guiC *appStruct.GuiComponent,cancel chan bool,timeExceed chan bool ){
	for {

		select {
		case <-timeExceed:
			return
		case <-time.After(time.Millisecond*50):
			if guiC.SearchIsActive == false{

				cancel <- false

				return
			}
		}

	}
}

func skipUnpack(guiC *appStruct.GuiComponent,skip chan bool,timeExceed chan bool){

	for {


		select {
		case <-timeExceed:
			return
		case <-time.After(time.Millisecond*50):
			if guiC.SkipItem == true{
				guiC.SkipItemNonArch = false
				guiC.SkipItem = false
				skip <- false
				return
			}
		}

	}
}


func wrapperUnpackCtx(path,ext,beautyName string,guiC *appStruct.GuiComponent,
	fi chan FrequencyInfo,fe chan ArchInfoError,end chan bool,ctx context.Context){

	unpackCtx(path,ext,beautyName,guiC,fi,fe,end,ctx)

	end <- true
}

func unpackCtx(path,ext,beautyName string,guiC *appStruct.GuiComponent,
	fi chan FrequencyInfo,fe chan ArchInfoError,end chan bool,ctx context.Context){

	var err error
	path = CheckExtension(path,ext)


	tempPath,err := tempDeleter.CreateTempDir("")
	guiC.AddTempDir <- tempPath

	defer func() {
		guiC.DeleteTempDir <-tempPath
	}()

	if beautyName == ""{
		beautyName = path
	}

	ok,err := IsSpaceEnough(path)

	if  !ok && err!= nil  && err.Error() =="недостаточно места для разархивации" {
		fe <- ArchInfoError{
			ArchiveName: beautyName,
			OpenError:   err,
		}
		return
	}


	if ext == ".7z" || ext == ".gz"{
		if check7z(path){

			fe <- ArchInfoError{
				ArchiveName: beautyName,
				OpenError:   errors.New("на файле пароль"),
			}

			return
		}
		err = unpackGZ(path,tempPath,ctx)
	} else {
		err = archiver.Unarchive(path, tempPath)

		if ext == ".zip" {
			RenameZipIncorrectName(tempPath)
		}

	}

	if !guiC.SearchIsActive {
		return
	}

	if err != nil {


		if checkForPassword(path,ext,err){
			//end <- true
			fe <- ArchInfoError{
				ArchiveName: beautyName,
				OpenError:   errors.New("на файле пароль"),
			}
			return
		}

		fe <- ArchInfoError{
			ArchiveName: beautyName,
			OpenError:   errors.New("На архиве пароль или он поврежден "),
		}
		return
	}

	findAnotherArhcWithCtx(tempPath,ext,beautyName ,guiC, fi,fe,end,ctx)

}

func findAnotherArhcWithCtx(path,ext,beautyName string,guiC *appStruct.GuiComponent,
	fi chan FrequencyInfo,fe chan ArchInfoError,end chan bool,ctx context.Context) {
	files, _ := ioutil.ReadDir(path)

	for _, f := range files {
		if !guiC.SearchIsActive {
			end <- true
			return
		}
		if f.IsDir() {
			//затенение делается чтобы имена папок не накладывались друг на друга
			//т.е. чтобы не было ситуаций вроде этой
			//в папке 1 лежат папки 2 и 3
			//и beatyName будет в конце выглядеть так 1/2/3 а не 1/2 и 1/3 т.к. 2 никуда не пропадет на след цикле
			beautyName := filepath.Join(beautyName, f.Name())
			findAnotherArhcWithCtx(filepath.Join(path, f.Name()), ext, beautyName, guiC,fi,fe,end,ctx)
		} else {
			ext_m,_ := mimetype.DetectFile(filepath.Join(path,f.Name()))
			ext = ext_m.Extension()
			if ext == ""{
				ae := ArchInfoError{
					ArchiveName: filepath.Join(beautyName,f.Name()),
					OpenError:   errors.New("Неизвестное расширение"),
					Ext: "",
				}
				fe <- ae
			}
			if isArch(ext){
				beautyName := filepath.Join(beautyName, f.Name())
				unpackCtx(filepath.Join(path,f.Name()),ext,beautyName,guiC,fi,fe,end,ctx)
				continue
			}
			w := textSearchAndExtract.FindText(filepath.Join(path,f.Name()),
				ext,newWordsConfig.GetDictWords(),guiC)
			fi <- FrequencyInfo{filepath.Join(beautyName, f.Name()),w,ext}
		}
	}
}



//func Unpack(path,ext,beautyName string,guiC *appStruct.GuiComponent)([]FrequencyInfo,[]ArchInfoError){
//	var archError []ArchInfoError
//	tempPath := filepath.Join(filepath.Dir(path), "temp_Path")
//	defer os.RemoveAll(tempPath)
//
//	if beautyName == ""{
//		beautyName = filepath.Join(filepath.Dir(path), filepath.Base(path))
//	}
//
//	err := archiver.Unarchive(path, tempPath)
//	if !guiC.SearchIsActive {
//		return nil,nil
//	}
//
//	if err != nil {
//		return nil,append(archError,ArchInfoError{
//			ArchiveName: beautyName,
//			OpenError:   err,
//		})
//	}
//
//	return findAnotherArhc(tempPath,ext,beautyName,guiC)
//}
//
//func findAnotherArhc(path,ext, beautyName string,guiC *appStruct.GuiComponent)([]FrequencyInfo,[]ArchInfoError){
//	var archError []ArchInfoError
//	var freqInfo []FrequencyInfo
//	files, _ := ioutil.ReadDir(path)
//	for _,f := range files{
//		if !guiC.SearchIsActive {
//			return nil,nil
//		}
//		if f.IsDir(){
//			beautyName = filepath.Join(beautyName, f.Name())
//			fi,ae := findAnotherArhc(filepath.Join(path,f.Name()),ext,beautyName,guiC)
//			freqInfo = append(freqInfo,fi...)
//			archError = append(archError,ae...)
//		} else {
//			ext_m,_ := mimetype.DetectFile(filepath.Join(path,f.Name()))
//			ext = ext_m.Extension()
//			if isArch(ext){
//				beautyName = filepath.Join(beautyName, f.Name())
//				fi,ae := Unpack(filepath.Join(path,f.Name()),ext,beautyName,guiC)
//				freqInfo = append(freqInfo,fi...)
//				archError = append(archError,ae...)
//				continue
//			}
//			w := textSearchAndExtract.FindText(filepath.Join(path,f.Name()),
//				ext,words.ReadWordsFromConfig())
//			freqInfo = append(freqInfo,FrequencyInfo{filepath.Join(beautyName, f.Name()),w})
//		}
//
//	}
//	return freqInfo,archError
//}















