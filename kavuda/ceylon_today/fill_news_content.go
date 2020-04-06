package ceylon_today

import (
	models3 "GIG-SDK/models"
	"GIG-SDK/request_handlers"
	models2 "GIG-Scripts/kavuda/ceylon_today/models"
	"GIG-Scripts/kavuda/models"
	"GIG-Scripts/kavuda/utils"
	"encoding/json"
)

func (d CeylonTodayDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	singleNewsResult, err := request_handlers.GetRequest(newsItem.Link)
	if err != nil {
		return models.NewsItem{}, "", err
	}

	var singleNewsResponse models2.SingleNewsResponse
	if err = json.Unmarshal([]byte(singleNewsResult), &singleNewsResponse); err != nil {
		return models.NewsItem{}, "", err
	}

	newsContent := singleNewsResponse.Data[0]

	newsItem = models.NewsItem{
		Title:   newsContent.Title,
		Snippet: newsContent.Snippet,
		Link:    newsItem.Link,
		Content: newsContent.HtmlContent,
		Date:    utils.ExtractPublishedDate("2006-01-02 15:04:05", newsContent.PublishDate),
		Author:  newsContent.AuthorName,
	}

	newsItem = models.UploadImagesToServer(newsItem, []models3.Upload{}, d.GetDefaultImageUrl())

	return newsItem, newsItem.Content, nil
}
