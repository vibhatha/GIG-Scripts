package ada_derana

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/kavuda/models"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
)

func (d AdaDeranaDecoder) FillNewsContent(newsItem extended_models.NewsArticle) (extended_models.NewsArticle, string, error) {
	return models.FillNewsContent(newsItem, ".news-content", "", clean_html.HtmlCleaner{}, d)
}
