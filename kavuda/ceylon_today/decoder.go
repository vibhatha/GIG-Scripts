package ceylon_today

import (
	"GIG-Scripts/kavuda/models"
)

var singleNewsUrl = "https://ceylontoday.lk/news/"
var defaultImageUrl = "https://ceylontoday.lk/images/logo.gif"
var newsSources = []models.NewsSource{
	{Link: "https://ceylontoday.lk/api/category/news?category_id=1&offset=0&time=all", Categories: []string{"Local"}},
	{Link: "https://ceylontoday.lk/api/category/news?category_id=3&offset=0&time=all", Categories: []string{"Business"}},
	{Link: "https://ceylontoday.lk/api/category/news?category_id=4&offset=0&time=all", Categories: []string{"Politics"}},
	{Link: "https://ceylontoday.lk/api/category/news?category_id=5&offset=0&time=all", Categories: []string{"Sports"}},
	{Link: "https://ceylontoday.lk/api/category/news?category_id=6&offset=0&time=all", Categories: []string{"Interviews"}},
}

type CeylonTodayDecoder struct {
	models.IDecoder
}

func (d CeylonTodayDecoder) GetSourceTitle() string {
	return "Ceylon Today"
}

func (d CeylonTodayDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
