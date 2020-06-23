package ceylon_today

import (
	"GIG-Scripts/kavuda/models"
)

var singleNewsUrl = "https://ceylontoday.lk/news/"
var newsSiteUrl = "https://ceylontoday.lk/api/category/news?category_id=1&offset=8&time=all"
var defaultImageUrl = "https://ceylontoday.lk/images/logo.gif"

type CeylonTodayDecoder struct {
	models.IDecoder
}

func (d CeylonTodayDecoder) GetSourceTitle() string {
	return "Ceylon Today"
}

func (d CeylonTodayDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
