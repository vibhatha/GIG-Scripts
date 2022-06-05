package daily_mirror

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/kavuda/models"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
)

func (d DailyMirrorDecoder) FillNewsContent(newsItem extended_models.NewsArticle) (extended_models.NewsArticle, string, error) {
	return models.FillNewsContent(newsItem, ".inner-content", "", clean_html.HtmlCleaner{
		Config: clean_html.Config{
			IgnoreElements: []string{"hr"},
		},
	}, d)
}
