package ceylon_today

import (
	"GIG-SDK/libraries/clean_html"
	"GIG-Scripts/kavuda/models"
)

func (d CeylonTodayDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".news-content", ".news-single-img", clean_html.HtmlCleaner{
		Config: clean_html.Config{
		}}, d)
}
