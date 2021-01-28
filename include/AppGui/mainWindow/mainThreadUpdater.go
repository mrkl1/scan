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
			UpdateHelper.RunUpdate(func() {
				mu.Lock()
				guiC.FileProgress.SetValue(count)
				defer mu.Unlock()
			})
		case count := <-guiC.InfoAboutScanningFilesUpdate:
			UpdateHelper.RunUpdate(func() {
				mu.Lock()
				defer mu.Unlock()
				guiC.InfoAboutScanningFiles.SetText(count)
				guiC.InfoAboutScanningFiles.AdjustSize()

			})
		case timeStr := <-guiC.ScanningTimeInfoUpdate:
			UpdateHelper.RunUpdate(func() {
				mu.Lock()
				defer mu.Unlock()
				guiC.ScanningTimeInfo.SetText(timeStr)
				guiC.ScanningTimeInfo.AdjustSize()

			})
		case updName := <-guiC.InfoAboutScanningFilesUpdate:
			UpdateHelper.RunUpdate(func() {
				mu.Lock()
				defer mu.Unlock()
				guiC.InfoAboutScanningFiles.SetText(updName)
				guiC.InfoAboutScanningFiles.AdjustSize()

			})
		case  <-guiC.FileTreeUpdate:
			UpdateHelper.RunUpdate(func() {

			})
		case  <-guiC.ErrorTableUpdate:
			UpdateHelper.RunUpdate(func() {

			})
		case  <-guiC.NonScanTableUpdate:
			UpdateHelper.RunUpdate(func() {

			})
		default:

			// TODO  попробовать сделать с qt timers
			// 200 - 500мс
			time.Sleep(150 * time.Millisecond)
		}

	}
}