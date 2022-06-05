package daily_mirror

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/libraries"
)

func (d DailyMirrorDecoder) ExtractNewsItems() ([]extended_models.NewsArticle, error) {
	var allNewsItems []extended_models.NewsArticle

	for _, newsSource := range newsSources {
		doc, err := global_helpers.GetDocumentFromUrl(newsSource.Link)
		if err != nil {
			return nil, err
		}

		allNewsItems = append(allNewsItems, createNewsItems(doc, newsSource)...)
	}

	return allNewsItems, nil
}

func createNewsItems(doc *goquery.Document, newsSource models.NewsSource) []extended_models.NewsArticle {
	var newsLinks []string
	newsNodes := doc.Find(".col-md-8")
	var newsItems []extended_models.NewsArticle
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)

		url, _ := nodeDoc.Find("a").First().Attr("href")
		if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
			newsLinks = append(newsLinks, url)

			newsItems = append(newsItems, createNewsItem(nodeDoc, url, newsSource))
		}
	}

	return newsItems
}

func createNewsItem(nodeDoc *goquery.Document, url string, newsSource models.NewsSource) extended_models.NewsArticle {
	dateString := nodeDoc.Find(".gtime").First().Nodes[0].FirstChild.Data
	title := nodeDoc.Find(".cat-hd-tx").First().Nodes[0].FirstChild.Data
	snippet := nodeDoc.Find("p").Last().Nodes[0].FirstChild.Data

	newsItem := *new(extended_models.NewsArticle).
		SetNewsTitle(title)
	newsItem.SetSource(url).
		SetSourceDate(helpers.ExtractPublishedDate("02 Jan 2006 ", dateString)).
		AddCategories(newsSource.Categories)
	newsItem.Snippet = snippet

	return newsItem
}
