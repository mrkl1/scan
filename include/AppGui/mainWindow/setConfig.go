package mainWindow

import (
	"github.com/myProj/scaner/new/include/appStruct"
	"github.com/myProj/scaner/new/include/config"
	"os"
)

//setConfig устанавливает стартовые параметры
//размеры окна, ширина столбцов, и тд
func setConfig(gc *appStruct.GuiComponent){
	cfg := config.GetCurrentConfig()
	sc := cfg.StartCoordinate
	gc.MainWindow.SetGeometry2(sc[0],sc[1],sc[2],sc[3])
	gc.StartDirectoryName = cfg.StartDir

	/*
	Нужно засунуть в отдельную функцию
	*/
	if cfg.StartDir == ""{
		gc.StartDirectoryForScan.SetText("Директория для сканирования не задана")
		gc.StartDirectoryForScan.AdjustSize()
		return
	}
	res,_ := exists(cfg.StartDir)
	if  res == false {
		// sudo chmod -R 600 testRoot/ например в таком случае
		gc.StartDirectoryForScan.SetText("Директории "+cfg.StartDir +" не существует или к ней невозможен доступ")
		gc.StartDirectoryForScan.AdjustSize()
		return
	}

	gc.StartDirectoryForScan.SetText(directoryForScanning+": "+cfg.StartDir)

	gc.StartDirectoryForScan.AdjustSize()
}

//проверка на существование папки или доступа к ней
func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}