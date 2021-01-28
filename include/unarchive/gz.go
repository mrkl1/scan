package unarchive

import (
	"bytes"
	"errors"
	"github.com/myProj/scaner/new/include/logggerScan"
	"os/exec"
	"runtime"
)

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

		logggerScan.SaveToLog("unpackGZ error"+stderr.String() )
		return errors.New("Отказ в доступе к файлу")
	}
	return err
}
