package main

import (
	"fmt"
	"github.com/mitchellh/panicwrap"
	"github.com/myProj/scaner/new/include/AppGui/mainWindow"
	"io/ioutil"
	"os"
	"time"
)

//messenger in golang
//https://github.com/easmith/p2p-messenger


func panicHandler(output string){
	ioutil.WriteFile("panic.logs",[]byte(time.Now().String()+"\n"+
		fmt.Sprintf("App panicked:\n\n%s\n", output)),0777)
	os.Exit(1)
}

func main(){
	exitStatus, err := panicwrap.BasicWrap(panicHandler)
	if err != nil {
		// Something went wrong setting up the panic wrapper. Unlikely,
		// but possible.
		panic(err)
	}


	if  exitStatus > -1{
		fmt.Println("exitStatus",exitStatus)
		os.Exit(exitStatus)
	}

	//if exitStatus >= 0 {
	//	os.Exit(exitStatus)
	//}

	println("app start")
	mainWindow.StartUI()
}

