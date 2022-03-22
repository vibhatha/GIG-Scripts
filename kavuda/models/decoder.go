package models

import (
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
	"github.com/lsflk/gig-sdk/models"
	"github.com/lsflk/gig-sdk/request_handlers"
	"golang.org/x/net/html"
	"strings"
)

type IDecoder interface {
	ExtractNewsItems() ([]NewsItem, error)
	FillNewsContent(newsItem NewsItem) (NewsItem, string, error)
	GetSourceTitle() string
	GetDefaultImageUrl() string
}

func FillNewsContent(newsItem NewsItem, contentClass string, imageClass string, htmlCleaner clean_html.HtmlCleaner, decoder IDecoder) (NewsItem, string, error) {
	if imageClass == "" {
		imageClass = contentClass
	}
	resp, err := request_handlers.GetRequest(newsItem.Link)
	if err != nil {
		return newsItem, "", err
	}

	newsDoc, err := libraries.HTMLStringToDoc(resp)
	if err != nil {
		return newsItem, "", err
	}

	newsSelection := newsDoc.Find(contentClass).First()
	newsHtml, err := newsSelection.Html()
	if err != nil {
		return newsItem, "", err
	}

	news, err := html.Parse(strings.NewReader(newsHtml))
	if err != nil {
		return newsItem, "", err
	}

	//for images in separate div
	newsImageSelection := newsDoc.Find(imageClass).First()
	newsImageHtml, err := newsImageSelection.Html()
	if err != nil {
		return newsItem, "", err
	}
	newsImage, err := html.Parse(strings.NewReader(newsImageHtml))
	if err != nil {
		return newsItem, "", err
	}

	//clean html code by removing unwanted information
	var imageList []models.Upload
	var contentImageList []models.Upload
	var imageUrl string

	newsItem.Content, _, contentImageList, imageUrl = htmlCleaner.CleanHTML(newsItem.Link, news)
	_, _, imageList, newsItem.ImageURL = htmlCleaner.CleanHTML(newsItem.Link, newsImage)
	imageList = append(imageList, contentImageList...)
	if newsItem.ImageURL == "" {
		newsItem.ImageURL = imageUrl
	}
	return UploadImagesToServer(newsItem, imageList, decoder.GetDefaultImageUrl()), newsSelection.Text(), nil
}

func UploadImagesToServer(newsItem NewsItem, imageList []models.Upload, defaultImageUrl string) NewsItem {
	if newsItem.ImageURL == "" {
		var imageUploadClass models.Upload
		newsItem.ImageURL, imageUploadClass = clean_html.GenerateImagePath(defaultImageUrl, defaultImageUrl)
		imageList = append(imageList, imageUploadClass)
	}

	for _, image := range imageList {
		go func(payload models.Upload) {
			request_handlers.UploadImage(payload)
		}(image)
	}

	return newsItem
}
