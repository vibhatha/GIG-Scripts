package daily_news

import (
	"GIG-SDK/libraries/clean_html"
	"Kavuda/models"
)

func (d DailyNewsDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".node-main-content .content", clean_html.HtmlCleaner{
		Config: clean_html.Config{
			IgnoreElements: []string{"section"},
			IgnoreClasses:  []string{"field-name-field-articletags", "print-edition", "field-name-field-section"},
		}}, d)
}
