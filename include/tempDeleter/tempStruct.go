package tempDeleter

import (
	"log"
	"os"
)

type folderDeleter struct {
	tempName string
	deleteFolder bool
}

func newFolder(name string)folderDeleter{
		return folderDeleter{
			tempName:     name,
			deleteFolder: false,
		}
}

func setFolderDelete(name string,fd []folderDeleter)[]folderDeleter{

	for i := 0;i < len (fd);i++ {
		if name == fd[i].tempName {
			fd[i].deleteFolder = true
			break
		}
	}
	return fd
}

func setAllTrue(fd []folderDeleter)[]folderDeleter{
	for i := 0;i < len (fd);i++ {
		fd[i].deleteFolder = true
	}
	return fd
}

func removeFolder(name string, fd []folderDeleter)[]folderDeleter{
	for i := 0;i < len (fd);i++ {
		if name == fd[i].tempName {
			ex,_ := exists(fd[i].tempName)
			if !ex {
				fd = append(fd[:i],fd[i+1:]...)
				break
			}

		}

	}
	return fd
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}


func removeAllDirs(fd []folderDeleter)[]folderDeleter{
	for i := 0; i < len(fd); i++ {
		err := os.RemoveAll(fd[i].tempName)
		if err != nil {
			log.Println("err rem all::",err)
			continue
		}
		log.Println("rem all::",fd[i].tempName)
		fd = removeFolder(fd[i].tempName,fd)
	}
	return fd
}