package daily_mirror

import (
	"GIG-Scripts/kavuda/models"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
)

func (d DailyMirrorDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".inner-content","", clean_html.HtmlCleaner{
		Config: clean_html.Config{
			IgnoreElements: []string{"hr"},
		},
	}, d)
}
