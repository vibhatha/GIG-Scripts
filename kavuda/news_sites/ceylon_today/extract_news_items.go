package ceylon_today

import (
	"GIG-Scripts"
	"GIG-Scripts/extended_models"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/news_sites/ceylon_today/models"
	"encoding/json"
	"github.com/lsflk/gig-sdk/libraries"
)

func (d CeylonTodayDecoder) ExtractNewsItems() ([]extended_models.NewsArticle, error) {
	var allNewsItems []extended_models.NewsArticle

	for _, newsSource := range newsSources {
		//get the page
		resp, err := GIG_Scripts.GigClient.GetRequest(newsSource.Link)
		if err != nil {
			return nil, err
		}
		var (
			newsItemsResponse models.NewsItemsResponse
			newsItems         []extended_models.NewsArticle
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
				newsItem := new(extended_models.NewsArticle).
					SetAuthor(newsItemResponse.Author).
					SetNewsTitle(newsItemResponse.Title)
				newsItem.SetSource(url).
					SetSourceDate(helpers.ExtractPublishedDate("2006-01-02 15:04:05", newsItemResponse.PublishDate)).
					AddCategories(newsSource.Categories)
				newsItem.Snippet = newsItemResponse.ShortContent
				newsItems = append(newsItems, *newsItem)
			}
		}
		allNewsItems = append(allNewsItems, newsItems...)
	}

	return allNewsItems, nil
}
