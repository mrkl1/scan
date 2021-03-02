package unarchive

import (
	"bytes"
	"context"
	"errors"
	"github.com/myProj/scaner/new/include/logggerScan"
	"log"
	"os/exec"
)

//https://stackoverflow.com/questions/11886531/terminating-a-process-started-with-os-exec-in-golang

func unpackGZ(path,dest string, ctx context.Context)error{

	var cmd *exec.Cmd
	//заменить на функцию, которая компилируется в зависимости от ОС

		cmd = getCommandContext(ctx,path,dest)


	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil && err.Error() != "signal: killed" {
		log.Println("unpackGZ/7z error :::"+stderr.String(),err)
		logggerScan.SaveToLog("unpackGZ/7z error"+stderr.String()+" "+err.Error() )
		return errors.New("Отказ в доступе к файлу")
	}
	return err
}
