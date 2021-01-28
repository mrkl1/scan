package newWordsConfig

import (
	"fmt"
	"github.com/myProj/scaner/new/include/logggerScan"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
)

func WordExport(){


	var fileDialog = widgets.NewQFileDialog2(nil,"Save as...",
		"","")
	fileDialog.SetMimeTypeFilters([]string{"*.txt","*.json"})
	newPAth := fileDialog.GetSaveFileName(nil,"Save as...",
		"","Text (*.txt);;Files (*.json)","",0)

	fmt.Println(newPAth,":",fileDialog.SelectedMimeTypeFilter(),fileDialog.SelectedNameFilter())
	d,err := ioutil.ReadFile(dictionaryFilename)
	if err != nil  {
		logggerScan.SaveToLog("WordExport ReadFile error: "+err.Error())
	}
	err = ioutil.WriteFile(newPAth,d,0666)
	if err != nil  {
		logggerScan.SaveToLog("WordExport WriteFile error: "+err.Error())
	}

}