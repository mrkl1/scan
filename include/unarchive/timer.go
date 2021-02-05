package unarchive

import (
	"fmt"
	"github.com/myProj/scaner/new/include/appStruct"
	"time"
)

var counter = 1

func setTimeEverySecond(guiC *appStruct.GuiComponent,st chan bool) {
	defer doneGor("setTimeEverySecond")
	//set zero time

	t := time.Time{}

	for {

		select {
		case <-st:
					guiC.ScanningTimeInfo.UpdateTextFromGoroutine("")
					guiC.ScanningTimeInfo.AdjustSizeFromGoroutine()
			return

		case <-time.After(1*time.Second):

			counter++
			t = t.Add(1000 * time.Millisecond)
			timeStr := fmt.Sprintf("Время сканирования файла: %02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
			guiC.ScanningTimeInfo.UpdateTextFromGoroutine(timeStr)
			guiC.ScanningTimeInfo.AdjustSizeFromGoroutine()

		}
	}
}