package extract

import (
	"code.sajari.com/docconv"
	"log"
)

func ExtractTextFromPdf(path string)string{

	res, err := docconv.ConvertPath(path)
	if err != nil {
		log.Println(err)
		return ""

	}

	return res.Body
}