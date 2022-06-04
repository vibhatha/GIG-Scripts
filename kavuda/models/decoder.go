package models

import (
	"GIG-Scripts"
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
	"github.com/lsflk/gig-sdk/models"
	"golang.org/x/net/html"
	"strings"
)

type IDecoder interface {
	ExtractNewsItems() ([]extended_models.NewsArticle, error)
	FillNewsContent(newsItem extended_models.NewsArticle) (extended_models.NewsArticle, string, error)
	GetSourceTitle() string
	GetDefaultImageUrl() string
}

func FillNewsContent(newsItem extended_models.NewsArticle, contentClass string, imageClass string, htmlCleaner clean_html.HtmlCleaner, decoder IDecoder) (extended_models.NewsArticle, string, error) {
	if imageClass == "" {
		imageClass = contentClass
	}
	resp, err := GIG_Scripts.GigClient.GetRequest(newsItem.Source)
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

	content, _, contentImageList, imageUrl := htmlCleaner.CleanHTML(newsItem.Source, news)
	newsItem = newsItem.SetContent(content)
	_, _, imageList, newsItem.ImageURL = htmlCleaner.CleanHTML(newsItem.Source, newsImage)
	imageList = append(imageList, contentImageList...)
	if newsItem.ImageURL == "" {
		newsItem.ImageURL = imageUrl
	}
	return UploadImagesToServer(newsItem, imageList, decoder.GetDefaultImageUrl()), newsSelection.Text(), nil
}

func UploadImagesToServer(newsItem extended_models.NewsArticle, imageList []models.Upload, defaultImageUrl string) extended_models.NewsArticle {
	if newsItem.ImageURL == "" {
		var imageUploadClass models.Upload
		newsItem.ImageURL, imageUploadClass = clean_html.GenerateImagePath(defaultImageUrl, defaultImageUrl)
		imageList = append(imageList, imageUploadClass)
	}

	for _, image := range imageList {
		go func(payload models.Upload) {
			GIG_Scripts.GigClient.UploadFile(payload)
		}(image)
	}

	return newsItem
}
