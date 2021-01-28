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
	guiC.InfoAboutScanningFiles.SetFixedHeight(guiC.InfoAboutScanningFiles.Height())
	t := time.Time{}

	for {

		select {
		case <-st:
			guiC.ScanningTimeInfoUpdate <-""
			return

		default:
			counter++

			timeStr := fmt.Sprintf("Время сканирования файла: %02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
			guiC.ScanningTimeInfoUpdate <-timeStr
			t = t.Add(1000 * time.Millisecond)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}