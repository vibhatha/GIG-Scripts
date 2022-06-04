package ceylon_today

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/kavuda/models"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
)

func (d CeylonTodayDecoder) FillNewsContent(newsItem extended_models.NewsArticle) (extended_models.NewsArticle, string, error) {
	return models.FillNewsContent(newsItem, ".news-content", ".news-single-img", clean_html.HtmlCleaner{
		Config: clean_html.Config{}}, d)
}
