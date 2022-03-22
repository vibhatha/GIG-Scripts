package helpers

import (
	models2 "GIG-Scripts/kavuda/models"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

/*
Decode NewsItem to entity
 */
func EntityFromNews(newsItem models2.NewsItem, category string) models.Entity {
	return models.Entity{
		Title:     newsItem.Title,
		Source: newsItem.Link,
		SourceDate: newsItem.Date,
		UpdatedAt: time.Now(),
		ImageURL:  newsItem.ImageURL,
		Snippet:   newsItem.Snippet,
	}.SetAttribute("content", models.Value{
		ValueType:     "html",
		ValueString: newsItem.Content,
	}).SetAttribute("date", models.Value{
		ValueType:     "date",
		ValueString: newsItem.Date.String(),
	}).SetAttribute("author", models.Value{
		ValueType:     "string",
		ValueString: newsItem.Author,
	}).AddCategory("News").AddCategory(category)

}
