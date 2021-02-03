//+build windows

package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/therecipe/qt/core"
	"sync"
	"time"
)

type updateHelper struct {
	core.QObject

	_ func(f func()) `signal:"runUpdate,auto`
}


func (*updateHelper) runUpdate(f func()) { f() }

var UpdateHelper = NewUpdateHelper(nil)

/*
нужно именно так обновлять компоненты т.к.
при таком подходе обновление компонентов происходит в главном потоке
если обновление компонентов производить не в главном потоке то
можно словить SIGSEGV: и приложение закрашится
 */
func runUpdater(guiC *appStruct.GuiComponent) {
	//TODO
	// возможно стоит сделать индикатор который показывает число просканированных
	// файлов из всех
	// пакеты с тегом для деплоя и с тегом для тестирования
	var mu sync.Mutex
	for {
		select {
		case  count := <-guiC.FileProgressUpdate:

			mu.Lock()
			guiC.FileProgress.SetValue(count)
			mu.Unlock()
			//WARNING проблема с этой функцией
			//UpdateHelper.RunUpdate(func() {
			//
			//})
		case count := <-guiC.InfoAboutScanningFilesUpdate:

			mu.Lock()
			guiC.InfoAboutScanningFiles.SetText(count)
			guiC.InfoAboutScanningFiles.AdjustSize()
			mu.Unlock()
			//UpdateHelper.RunUpdate(func() {
			//
			//})
		case timeStr := <-guiC.ScanningTimeInfoUpdate:
			mu.Lock()
			guiC.ScanningTimeInfo.SetText(timeStr)
			guiC.ScanningTimeInfo.AdjustSize()
			mu.Unlock()

			//UpdateHelper.RunUpdate(func() {
			//})

		case updName := <-guiC.InfoAboutScanningFilesUpdate:
			mu.Lock()
			guiC.InfoAboutScanningFiles.SetText(updName)
			guiC.InfoAboutScanningFiles.AdjustSize()
			mu.Unlock()

			//UpdateHelper.RunUpdate(func() {
			//})
		//case  <-guiC.FileTreeUpdate:
		//	UpdateHelper.RunUpdate(func() {
		//
		//	})
		//case  <-guiC.ErrorTableUpdate:
		//	UpdateHelper.RunUpdate(func() {
		//
		//	})
		//case  <-guiC.NonScanTableUpdate:
		//	UpdateHelper.RunUpdate(func() {
		//
		//	})
		default:

			// TODO  попробовать сделать с qt timers
			// 200 - 500мс
			time.Sleep(150 * time.Millisecond)
		}

	}
}