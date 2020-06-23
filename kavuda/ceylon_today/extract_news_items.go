package ceylon_today

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	models2 "GIG-Scripts/kavuda/ceylon_today/models"
	"GIG-Scripts/kavuda/models"
	"GIG-Scripts/kavuda/utils"
	"encoding/json"
)

func (d CeylonTodayDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}
	var (
		newsItemsResponse models2.NewsItemsResponse
		newsItems         []models.NewsItem
	)
	if err := json.Unmarshal([]byte(resp), &newsItemsResponse); err != nil {
		return nil, err
	}

	var newsLinks []string

	//create news item list from news item responses
	for _, newsItemResponse := range newsItemsResponse.Data {
		url := singleNewsUrl + newsItemResponse.Slug
		if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
			newsLinks = append(newsLinks, url)
			newsItem := models.NewsItem{
				Link: url,
				Title:newsItemResponse.Title,
				Snippet:newsItemResponse.ShortContent,
				Date:utils.ExtractPublishedDate("2006-01-02 15:04:05", newsItemResponse.PublishDate),
				Author:newsItemResponse.Author,

			}
			newsItems = append(newsItems, newsItem)
		}
	}

	return newsItems, nil
}
