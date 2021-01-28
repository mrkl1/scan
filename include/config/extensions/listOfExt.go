package extensions

//TODO В этот пакет добавить проверки на архивы
// видео аудио и тд
// +генерить при необходимости деф файл
// попробовать добавить через канаl file tree


var IncludeExtension = []string{
	//text
	".txt",".json",".xml",
	".html",".py",".js",".php",".pl",".lua",".tcl",
	//old office
	".doc",".ppt",".vsd",".xls",
	//new office
	".docx",".pptx",".vsdx",".xlsx",
	//other document
	".pdf",".rtf",
	//archives
	".tar",".tar.gz",".zip",".rar",".7z",".gzip",".tar.xz",
	".xz",".gz",
	//images
	".png",".jpg",".jpeg",".ico",".bmp","tiff",
	//videos
	//25 video formats
	".mov",".avi",".flv",".mkv",".mpeg",".asf",".mp4",".3gp",
	".f4v",".hevc",".m2ts",".m2v",".m4v",".mjpeg",".mpg",".mts",
	".mxf",".ogv",".rm",".swf",".ts",".vob",".webm",".wmv",".wtv",
	//music
	".mp3",".flac",".aac",".waw","wma",
	//одно и тоже расширение
	".ogg", ".ogv", ".oga", ".ogx", ".spx", ".opus,",".ogm",
	".wav",".wma",
}

var UnsupportedExtension = []string{
	".ppt",".vsd",
}

var Archives = []string{
	".tar",".tar.gz",".zip",".rar",".7z",".gzip",".tar.xz",
	".xz",".gz",
}

//тут много примеров
//https://filesamples.com/categories/video
var ImagesVideoMusic = []string{
	//images
	".png",".jpg",".jpeg",".ico",".bmp","tiff",
	//videos
	//25 video formats
	".mov",".avi",".flv",".mkv",".mpeg",".asf",".mp4",".3gp",
	".f4v",".hevc",".m2ts",".m2v",".m4v",".mjpeg",".mpg",".mts",
	".mxf",".ogv",".rm",".swf",".ts",".vob",".webm",".wmv",".wtv",
	//music
	".mp3",".flac",".aac",".waw","wma",
	".ogg", ".ogv", ".oga", ".ogx", ".spx", ".opus,",".ogm",
	".wav",".wma",
}
