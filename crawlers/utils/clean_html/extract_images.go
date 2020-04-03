package clean_html

import (
	"GIG/app/models"
	"GIG/commons"
	"golang.org/x/net/html"
	"strconv"
)

func ExtractImages(startTag string, n *html.Node, uri string, imageList []models.Upload) (string, []models.Upload, string, int) {
	sourceLink := ""
	var uploadImageClass models.Upload
	imageWidth := 0

	if n.Data == "img" {
		var (
			src   html.Attribute
			width html.Attribute
		)
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				src = attr
			} else if attr.Key == "width" {
				width = attr
			}
		}

		sourceLink, uploadImageClass = GenerateImagePath(src.Val, uri)
		imageWidth, _ = strconv.Atoi(width.Val)
		startTag = n.Data + " src='" + sourceLink+"'"
		imageList = append(imageList, uploadImageClass)
	}
	return startTag, imageList, sourceLink, imageWidth
}

func GenerateImagePath(href string, uri string) (string, models.Upload) {
	fixedSrc := libraries.FixUrl(href, uri)
	fileName := libraries.ExtractFileName(fixedSrc)
	bucketName := libraries.ExtractDomain(fixedSrc)
	return "images/" + bucketName + "/" + fileName, models.Upload{Title: bucketName, Source: fixedSrc}
}
