package unarchive

import (
	"fmt"
	"github.com/myProj/scaner/new/include/appStruct"
	"sync"
	"time"
)

var counter = 1

func setTimeEverySecond(guiC *appStruct.GuiComponent,st chan bool) {

	//set zero time



	t := time.Time{}

	for {

		select {
		case <-st:
			guiC.ScanningTimeInfo.UpdateTextFromGoroutine("")
			guiC.ScanningTimeInfo.AdjustSizeFromGoroutine()
			var mu sync.Mutex
			mu.Lock()
			guiC.IsTimeUpdate = false
			mu.Unlock()
			return

		case <-time.After(1*time.Second):

			if guiC.IsTimeUpdate{

				counter++
				t = t.Add(1000 * time.Millisecond)
				timeStr := fmt.Sprintf("Время сканирования файла: %02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				guiC.ScanningTimeInfo.UpdateTextFromGoroutine(timeStr)
				guiC.ScanningTimeInfo.AdjustSizeFromGoroutine()
				//в общем иногда каналы
				//неправильно закрываются поэтому тут
				//нужен это код
				if !guiC.IsTimeUpdate{
					guiC.ScanningTimeInfo.UpdateTextFromGoroutine("")
					guiC.ScanningTimeInfo.AdjustSizeFromGoroutine()
					return
				}
			}

		}
	}
}