//+build linux

package unarchive

import (
	"context"
	"os/exec"
)

func getCommandContext(ctx context.Context,path,dest string)*exec.Cmd{
	return exec.CommandContext(ctx,pass7zLinux,"e",path,"-o"+dest)
}
