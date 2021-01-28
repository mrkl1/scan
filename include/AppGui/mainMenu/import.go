package mainMenu

import (
	"github.com/myProj/scaner/new/include/config/newWordsConfig"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/therecipe/qt/widgets"
	"os"
)

func importFile() {
	filename := widgets.QFileDialog_GetOpenFileName(nil,"Open Directory","","","",widgets.QFileDialog__DontUseNativeDialog)

	f, err := os.Open(filename)
	if err != nil {
		logggerScan.SaveToLog("importFile error"+err.Error())
		return
	}
	defer f.Close()
	newWordsConfig.MergeImports(filename)


}
