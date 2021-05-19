//+build windows

package unarchive

import (
	"context"
	"os/exec"
	"syscall"
)

func getCommandContext(ctx context.Context,path,dest string)*exec.Cmd{
	var cmd *exec.Cmd
	cmd = exec.CommandContext(ctx,pass7zWindows,"e",path,"-o"+dest)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

//check archive Size
func getCommandCheckSize(path string)*exec.Cmd{
	return exec.Command(pass7zWindows,"l",path)
}