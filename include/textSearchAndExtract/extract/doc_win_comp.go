//+build windows

package extract

func DocToTxt(filePath string)string{
	var out bytes.Buffer
	var stderr bytes.Buffer
	//os.Setenv("FOO", "1")
	fp,_ :=  filepath.Abs("anti/bin/antiword.exe")

	os.Setenv("HOME", fp)
	cmd := exec.Command("anti\\bin\\antiword.exe","-m" ,"UTF-8.txt","-t",filePath)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//&syscall.SysProcAttr{CreationFlags: 0x08000000} //вообще не создавать консольное окно
	//cmd := exec.Command(".\\bin\\antiword.exe","-t",filePath)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	//err
	cmd.Run()
	//fmt.Println("antiword error",fmt.Sprint(err) + ": " + stderr.String())
	return out.String()

}

//linux
