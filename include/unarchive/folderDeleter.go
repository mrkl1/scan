package unarchive

/*

	в канал приходят имена директорий со временными папками в структуру, в которой указано, что
	эту папку пока не нужно удалять
	После канал опять принимает структуру, в которой говорится что нужно удалить эту папку,

	они добавляются

	возможно стоить логировать

	также написать генератор кода, который сможет создавать
	логи в местах где стоит комментарий с LOG



func main() {
	//tFile, err := ioutil.TempFile("", "gostdcookbook")
	//if err != nil {
	//	panic(err)
	//}
	//// Вызов ответственен за очистку
	////defer os.Remove(tFile.Name())
	//
	//fmt.Println(tFile.Name())
	//fmt.Println( filepath.Abs(tFile.Name()))
	//
	//// TempDir возвращает путь
	//tDir, err := ioutil.TempDir("", "gostdcookbookdir")
	//fmt.Println( filepath.Abs(tDir))
	//if err != nil {
	//	panic(err)
	//}
	////defer os.Remove(tDir)
	//fmt.Println(tDir)

	fmt.Println(exists("/tmp/gostdcookbook225651541"))
	fmt.Println(exists("/tmp/gostdcookbookdir165217456"))
	os.Remove("/tmp/gostdcookbook225651541")
	os.Remove("/tmp/gostdcookbookdir165217456")

}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}
 */

//func StartDelete(folderName chan string){
//
//	for {
//
//	}
//
//}
//





