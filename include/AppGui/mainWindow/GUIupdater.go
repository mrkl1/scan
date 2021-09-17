package mainWindow

import (
	"fmt"
	"github.com/myProj/scaner/new/include/appStruct"
	"time"
)


func GUIUpdater(guiC *appStruct.GuiComponent){


	for {
		select {

		case <-guiC.EndUIUpdate:
			fmt.Println("STOP GUI upd")
			guiC.FileProgress.ValueChangedFromGoroutine(guiC.ProgressBarValue)
			return
		case <-time.After(400*time.Millisecond):
			guiC.FileProgress.ValueChangedFromGoroutine(guiC.ProgressBarValue)

		}
	}
}

