package ada_derana

import (
	"GIG-SDK/libraries/clean_html"
	"Kavuda/models"
)

func (d AdaDeranaDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".news-content", clean_html.HtmlCleaner{
	},d)
}