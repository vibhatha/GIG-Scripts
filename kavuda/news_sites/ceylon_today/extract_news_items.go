package ceylon_today

import (
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	models2 "GIG-Scripts/kavuda/news_sites/ceylon_today/models"
	"encoding/json"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/request_handlers"
)

func (d CeylonTodayDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	var allNewsItems []models.NewsItem

	for _, newsSource := range newsSources {
		//get the page
		resp, err := request_handlers.GetRequest(newsSource.Link)
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
					Link:       url,
					Title:      newsItemResponse.Title,
					Snippet:    newsItemResponse.ShortContent,
					Date:       helpers.ExtractPublishedDate("2006-01-02 15:04:05", newsItemResponse.PublishDate),
					Author:     newsItemResponse.Author,
					Categories: newsSource.Categories,
				}
				newsItems = append(newsItems, newsItem)
			}
		}
		allNewsItems = append(allNewsItems, newsItems...)
	}

	return allNewsItems, nil
}
