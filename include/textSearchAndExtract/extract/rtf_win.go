//+build windows

package extract

import (
	"bytes"
	"fmt"
	"github.com/k3a/html2text"
	"log"
	"os/exec"
	"path/filepath"
	"syscall"
)

func Rtf2txt(docPath string)string{
	execUnrtfPath := filepath.Join("unrtf","bin","unrtf64.exe")

	arguments     := []string{"-P",filepath.Join("unrtf","share"),"--html",docPath}
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(execUnrtfPath,arguments...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("unrtf error:"+fmt.Sprint(err) + ": " + stderr.String())
		return ""
	}
return  html2text.HTML2Text(out.String())
}
