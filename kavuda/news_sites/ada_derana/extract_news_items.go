package ada_derana

import (
	"GIG-SDK/libraries"
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func (d AdaDeranaDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	var allNewsItems []models.NewsItem

	for _, newsSource := range newsSources {
		doc, err := global_helpers.GetDocumentFromUrl(newsSource.Link)
		if err != nil {
			return nil, err
		}

		var newsLinks []string

		newsNodes := doc.Find(".news-story")
		var newsItems []models.NewsItem
		for _, node := range newsNodes.Nodes {
			nodeDoc := goquery.NewDocumentFromNode(node)

			url, _ := nodeDoc.Find("a").First().Attr("href")

			if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
				newsLinks = append(newsLinks, url)

				title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
				snippet := nodeDoc.Find("p").First().Nodes[0].FirstChild.Data
				dateString := strings.Replace(nodeDoc.Find("span").Last().Nodes[0].FirstChild.Data, "  ", " ", -1) // replacing double &nbsp; characters with space
				dateString = strings.Replace(dateString, "| ", "", -1)

				newsItems = append(newsItems, models.NewsItem{
					Title:      title,
					Link:       url,
					Date:       helpers.ExtractPublishedDate("January 2, 2006 3:04 pm", dateString),
					Snippet:    snippet,
					Categories: newsSource.Categories,
				})
			}
		}
		allNewsItems = append(allNewsItems, newsItems...)
	}

	return allNewsItems, nil
}