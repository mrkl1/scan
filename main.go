package main

import (
	"github.com/myProj/scaner/new/include/AppGui/mainWindow"
)

//messenger in golang
//https://github.com/easmith/p2p-messenger
//https://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang
func main(){

	//defer func() {
	//
	//		logggerScan.SaveToLog(string(debug.Stack()))
	//	}()
	mainWindow.StartUI()
}

