package daily_news

import (
	"GIG-Scripts/kavuda/models"
)

var newsSiteUrl1 = "http://www.dailynews.lk/category/local"
var newsSiteUrl2 = "http://www.dailynews.lk/category/political"
var newsSiteUrl3 = "http://www.dailynews.lk/category/business"
var defaultImageUrl = "http://www.dailynews.lk/sites/all/themes/dailynews/logo.png"

type DailyNewsDecoder struct {
	models.IDecoder
}

func (d DailyNewsDecoder) GetSourceTitle() string {
	return "Daily News"
}

func (d DailyNewsDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
