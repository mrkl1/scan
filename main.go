package main

import (
	"fmt"
	"github.com/mitchellh/panicwrap"
	"github.com/myProj/scaner/new/include/AppGui/mainWindow"
	"log"
	"os"
	"time"
)

//messenger in golang
//https://github.com/easmith/p2p-messenger


func panicHandler(output string){

	f,_ := os.OpenFile("panic.logs",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)

	log.SetOutput(f)
	log.Println(time.Now().String()+"\n"+
		fmt.Sprintf("App panicked:\n\n%s\n", output))
	log.SetOutput(os.Stdout)
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

