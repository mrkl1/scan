package searchFilter

import (
	"github.com/myProj/scaner/new/include/config/extensions"
	"strings"
)

//возможно тут стоит использовать бинарный поиск
func IsExtensionForSearch(extension string)bool{
	ae := extensions.GetAllowList()

	for _,ext := range ae{
		if ext.Ext == extension{
			return true
		}
	}
	return false
}

func IsArchive(extension string)bool{
	for _,ext := range extensions.Archives{
		if ext == extension{
			return true
		}
	}
	return false
}

func IsMedia(extension string)bool{
	for _,ext := range extensions.ImagesVideoMusic{
		if ext == extension{
			return true
		}
	}
	return false
}

func IsUnsupported(extension string)bool{
	for _,ext := range extensions.UnsupportedExtension{
		if ext == extension{
			return true
		}
	}
	return false
}



func IsContainArchiveExtension(file string)bool{
	for _,ext := range extensions.Archives {
		if strings.HasSuffix(file,ext){
			return true
		}
	}
	return false
}




