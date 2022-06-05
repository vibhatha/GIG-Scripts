package the_island

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/kavuda/models"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
)

func (d TheIslandDecoder) FillNewsContent(newsItem extended_models.NewsArticle) (extended_models.NewsArticle, string, error) {
	return models.FillNewsContent(newsItem, "#mvp-content-main", "#mvp-post-feat-img", clean_html.HtmlCleaner{
		Config: clean_html.Config{
			IgnoreElements: []string{"h1"},
			IgnoreStrings:  []string{"Add new comment", "Print Edition", "Send to Friend"},
			IgnoreClasses:  []string{"article_info_col", "article_date"},
		}}, d)
}
