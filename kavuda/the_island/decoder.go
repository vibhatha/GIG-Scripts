package the_island

import (
	"GIG-Scripts/kavuda/models"
)

var	defaultImageUrl="http://island.lk/userfiles/image/danweem/island.gif"
var newsSources = []models.NewsSource{
	{Link: "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=39", Categories: []string{"Local"}},
	{Link: "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=48", Categories: []string{"Local"}},
	{Link: "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=49", Categories: []string{"Features"}},
	{Link: "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=50", Categories: []string{"Sports"}},
	{Link: "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=51", Categories: []string{"Business"}},
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
