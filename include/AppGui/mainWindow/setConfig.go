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
		gc.StartDirectoryForScan.UpdateTextFromGoroutine("Директория для сканирования не задана")
		gc.StartDirectoryForScan.AdjustSizeFromGoroutine()
		return
	}
	res,_ := exists(cfg.StartDir)
	if  res == false {
		// sudo chmod -R 600 testRoot/ например в таком случае
		gc.StartDirectoryForScan.UpdateTextFromGoroutine("Директории "+cfg.StartDir +" не существует или к ней невозможен доступ")
		gc.StartDirectoryForScan.AdjustSizeFromGoroutine()
		return
	}

	gc.StartDirectoryForScan.UpdateTextFromGoroutine(directoryForScanning+": "+cfg.StartDir)

	gc.StartDirectoryForScan.AdjustSizeFromGoroutine()
}

//проверка на существование папки или доступа к ней
func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}