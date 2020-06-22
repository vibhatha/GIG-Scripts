package ceylon_today

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	models2 "GIG-Scripts/kavuda/ceylon_today/models"
	"GIG-Scripts/kavuda/models"
	"encoding/json"
	"fmt"
	"strconv"
)

func (d CeylonTodayDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
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
		fmt.Println(newsItemResponse)
		url := singleNewsUrl + strconv.Itoa(newsItemResponse.NewsId)
		if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
			newsLinks = append(newsLinks, url)
			//newsItem := models.NewsItem{
			//	Link: url,
			//}
			//newsItems = append(newsItems, newsItem)
		}
	}

	return newsItems, nil
}
