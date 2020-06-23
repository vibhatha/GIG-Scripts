package daily_news

import (
	"GIG-Scripts/kavuda/models"
)

var defaultImageUrl = "http://www.dailynews.lk/sites/all/themes/dailynews/logo.png"
var newsSources = []models.NewsSource{
	{Link: "http://www.dailynews.lk/category/local", Categories: []string{"Local"}},
	{Link: "http://www.dailynews.lk/category/political", Categories: []string{"Political"}},
	{Link: "http://www.dailynews.lk/category/business", Categories: []string{"Business"}},
}

type DailyNewsDecoder struct {
	models.IDecoder
}

func (d DailyNewsDecoder) GetSourceTitle() string {
	return "Daily News"
}

func (d DailyNewsDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
