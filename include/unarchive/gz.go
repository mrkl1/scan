package unarchive

import (
	"bytes"
	"errors"
	"github.com/myProj/scaner/new/include/logggerScan"
	"log"
	"os/exec"
	"runtime"
)

//https://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang

/*
TODO 3.01.2021 тут есть явная проблема
 основаная горутина прерывается но Run нет
 т.к. представляет собой отедельную горутину
 поэтому ее нужно завершать отдельно черещ
 cmd.Process.Kill

 как вариант сделать структуру на уровне пакета которая
 будет завершать все эти команды, при завершении работы
 основной функции разархивации

 */

func unpackGZ(path,dest string)error{

	var cmd *exec.Cmd

	if runtime.GOOS == "windows"{
		cmd = exec.Command(pass7zWindows,"e",path,"-o"+dest)
	} else if runtime.GOOS == "linux"{
		cmd = exec.Command(pass7zLinux,"e",path,"-o"+dest)
	} else {
		return errors.New("Не поддерживается для данной ОС")
	}



	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	err := cmd.Run()


	if err != nil {

		log.Println("unpackGZ/7z error"+stderr.String())
		logggerScan.SaveToLog("unpackGZ/7z error"+stderr.String() )
		return errors.New("Отказ в доступе к файлу")
	}
	return err
}
