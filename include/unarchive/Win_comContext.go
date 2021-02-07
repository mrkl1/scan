//+build windows

package unarchive

import (
	"os/exec"
	"context"
	"syscall"
)

func getCommandContext(ctx context.Context,path,dest string)*exec.Cmd{
	var cmd *exec.Cmd
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd = exec.CommandContext(ctx,pass7zWindows,"e",path,"-o"+dest)
	return cmd
}
