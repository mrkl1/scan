//+build linux

package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
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
		for {
		select {
		case  <-guiC.NonScanTableUpdate:

			UpdateHelper.RunUpdate(func() {
				widgets.QMessageBox_Information(nil, "Ошибка доступа к директории","Невозможно получить доступ к директории", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			})
		default:

			// TODO  попробовать сделать с qt timers
			// 200 - 500мс
			time.Sleep(150 * time.Millisecond)
		}

	}
}