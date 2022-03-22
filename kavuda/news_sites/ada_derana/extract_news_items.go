package ada_derana

import (
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/libraries"
	"strings"
)

func (d AdaDeranaDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	var allNewsItems []models.NewsItem

	for _, newsSource := range newsSources {
		doc, err := global_helpers.GetDocumentFromUrl(newsSource.Link)
		if err != nil {
			return nil, err
		}

		allNewsItems = append(allNewsItems, createNewsItems(doc,newsSource)...)
	}

	return allNewsItems, nil
}

func createNewsItems(doc *goquery.Document, newsSource models.NewsSource) []models.NewsItem {
	var newsLinks []string
	newsNodes := doc.Find(".news-story")
	var newsItems []models.NewsItem
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

func createNewsItem(nodeDoc *goquery.Document, url string, newsSource models.NewsSource) models.NewsItem {
	title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
	snippet := nodeDoc.Find("p").First().Nodes[0].FirstChild.Data
	dateString := strings.Replace(nodeDoc.Find("span").Last().Nodes[0].FirstChild.Data, "  ", " ", -1) // replacing double &nbsp; characters with space
	dateString = strings.Replace(dateString, "| ", "", -1)

	return models.NewsItem{
		Title:      title,
		Link:       url,
		Date:       helpers.ExtractPublishedDate("January 2, 2006 3:04 pm", dateString),
		Snippet:    snippet,
		Categories: newsSource.Categories,
	}
}
