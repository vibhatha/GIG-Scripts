package the_island

import (
	"GIG-Scripts/kavuda/models"
)

var	defaultImageUrl="http://island.lk/userfiles/image/danweem/island.gif"
var newsSources = []models.NewsSource{
	{Link: "https://island.lk/category/news/", Categories: []string{"Local"}},
	{Link: "https://island.lk/category/features/", Categories: []string{"Features"}},
	{Link: "https://island.lk/category/sports/", Categories: []string{"Sports"}},
	{Link: "https://island.lk/category/business/", Categories: []string{"Business"}},
	{Link: "https://island.lk/category/politics/", Categories: []string{"Politics"}},
}

type TheIslandDecoder struct {
	models.IDecoder
}

func (d TheIslandDecoder) GetSourceTitle() string {
	return "The Island"
}

func (d TheIslandDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
