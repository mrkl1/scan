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
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000,HideWindow: true}
	return cmd
}

//check archive Size
func getCommandCheckSize(path string)*exec.Cmd{
	cmd := exec.Command(pass7zWindows,"l",path)
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000,HideWindow: true}
	return cmd
}


func getCommandPassword(path string)*exec.Cmd{
	cmd := exec.Command(pass7zWindows,"t",path,"-p")
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000,HideWindow: true}
	return cmd
}
