package the_island

import (
	"Kavuda/models"
)

var newsSiteUrl = "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=39"
var newsSiteUrl1 = "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=48"
var	defaultImageUrl="http://island.lk/userfiles/image/danweem/island.gif"

type TheIslandDecoder struct {
	models.IDecoder
}

func (d TheIslandDecoder) GetSourceTitle() string {
	return "The Island"
}

func (d TheIslandDecoder) GetDefaultImageUrl() string {
	return defaultImageUrl
}
