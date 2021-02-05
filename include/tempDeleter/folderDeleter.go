package tempDeleter

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"io/ioutil"
	"log"
	"os"
	"time"
)


/*
WARNING еще нужно дополнительное тестирование
 */
func StartDelete(guiC *appStruct.GuiComponent ){

	var tempDirs []folderDeleter

	scanningEnd := false
	for {
		select {
		case fName := <-guiC.AddTempDir:
			tempDirs = append(tempDirs, newFolder(fName))
		case dfName := <-guiC.DeleteTempDir:

			tempDirs = setFolderDelete(dfName, tempDirs)
			err := os.RemoveAll(dfName)
			if err == nil {
				tempDirs = removeFolder(dfName,tempDirs)
			} else {
				log.Println(err)
			}
		case <-guiC.EndDeleteTemp:

			tempDirs = setAllTrue(tempDirs)
			//log.Println("EndDeleteTemp ::: ", tempDirs)
			scanningEnd = true
		case <-time.After(2000*time.Millisecond):


			if scanningEnd {

				tempDirs = removeAllDirs(tempDirs)

				if len(tempDirs) == 0 {
					//log.Println("Delete folder complete")
					return
				}
			}
		}
	}

}



func CreateTempDir(tempPath string)(string,error){
	tDir, err := ioutil.TempDir(tempPath, "tempArch")
	if err != nil {
		return "",err
	}
	return tDir,nil
}





