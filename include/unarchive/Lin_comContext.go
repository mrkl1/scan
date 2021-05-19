//+build linux

package unarchive

import (
	"context"
	"os/exec"
)

func getCommandContext(ctx context.Context,path,dest string)*exec.Cmd{
	return exec.CommandContext(ctx,pass7zLinux,"e",path,"-o"+dest)
}

//check archive Size
func getCommandCheckSize(path string)*exec.Cmd{
	return exec.Command(pass7zLinux,"l",path,"-p")
}

func getCommandPassword(path string)*exec.Cmd{
	return exec.Command(pass7zLinux,"t",path,"-p")
}