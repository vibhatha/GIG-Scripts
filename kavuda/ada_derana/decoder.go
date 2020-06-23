package ada_derana

import (
	"GIG-Scripts/kavuda/models"
)

var defaultImageUrl = "http://www.adaderana.lk/newstyle2017/assets/images/adaderana-header-logo-2020.png"
var newsSources = []models.NewsSource{
	{Link: "http://www.adaderana.lk/hot-news/", Categories: []string{"Local"}},
	{Link: "http://www.adaderana.lk/sports-news/", Categories: []string{"Sports"}},
}

type AdaDeranaDecoder struct {
	models.IDecoder
}

func (d AdaDeranaDecoder) GetSourceTitle() string {
	return "Ada Derana"
}

func (d AdaDeranaDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
