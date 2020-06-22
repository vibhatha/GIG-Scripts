package ceylon_today

import (
	"GIG-Scripts/kavuda/models"
)

var singleNewsUrl = "https://ceylontoday.lk/site-api/api/News/getSingleNews?Id="
//var newsSiteUrl = "https://ceylontoday.lk/site-api/api/News/getNews?Id&searchValue&orderColum&orderType&start=0&limit=20&newsCategoryId=8&timePeriod=ALL"
var newsSiteUrl = "https://ceylontoday.lk/api/category/news?category_id=1&offset=8&time=all"
var defaultImageUrl = "https://ceylontoday.lk/assets/images/header_logo.gif"

type CeylonTodayDecoder struct {
	models.IDecoder
}

func (d CeylonTodayDecoder) GetSourceTitle() string {
	return "Ceylon Today"
}

func (d CeylonTodayDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
