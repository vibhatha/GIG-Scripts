package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

type NewsArticle struct {
	models.Entity
}

func (n NewsArticle) SetNewsTitle(title string) NewsArticle {
	n.Entity = n.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: title,
	})
	return n
}

func (n NewsArticle) SetContent(content string) NewsArticle {
	n.Entity = n.SetAttribute("content", models.Value{
		ValueType:   ValueType.HTML,
		ValueString: content,
	})
	return n
}

func (n NewsArticle) SetDate(date time.Time) NewsArticle {
	n.Entity = n.SetAttribute("date", models.Value{
		ValueType:   ValueType.Date,
		ValueString: date.String(),
	})
	n.SourceDate = date
	return n
}

func (n NewsArticle) SetAuthor(author string) NewsArticle {
	n.Entity = n.SetAttribute("date", models.Value{
		ValueType:   ValueType.String,
		ValueString: author,
	})
	return n
}

//func EntityFromNews(newsItem models2.NewsItem, category string) models.Entity {
//	return models.Entity{
//		Title:      newsItem.Title,
//		Source:     newsItem.Link,
//		SourceDate: newsItem.Date,
//		UpdatedAt:  time.Now(),
//		ImageURL:   newsItem.ImageURL,
//		Snippet:    newsItem.Snippet,
//	}
//}).SetAttribute("date", models.Value{
//ValueType:     "date",
//ValueString: newsItem.Date.String(),
//}).SetAttribute("author", models.Value{
//ValueType:     "string",
//ValueString: newsItem.Author,
//}).AddCategory("News").AddCategory(category)
//
//}
