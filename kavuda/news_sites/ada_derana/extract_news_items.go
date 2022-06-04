package ada_derana

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/libraries"
	"strings"
)

func (d AdaDeranaDecoder) ExtractNewsItems() ([]extended_models.NewsArticle, error) {
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
	newsNodes := doc.Find(".news-story")
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
	title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
	snippet := nodeDoc.Find("p").First().Nodes[0].FirstChild.Data
	dateString := strings.Replace(nodeDoc.Find("span").Last().Nodes[0].FirstChild.Data, "  ", " ", -1) // replacing double &nbsp; characters with space
	dateString = strings.Replace(dateString, "| ", "", -1)

	newsItem := extended_models.NewsArticle{}.SetNewsTitle(title)
	newsItem.Source = url
	newsItem.Snippet = snippet
	newsItem = newsItem.SetDate(helpers.ExtractPublishedDate("January 2, 2006 3:04 pm", dateString))
	newsItem.Entity = newsItem.AddCategories(newsSource.Categories)
	return newsItem
}
