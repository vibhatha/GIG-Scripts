package ada_derana

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	utils2 "GIG-Scripts/wikipedia/utils"
	"Kavuda/models"
	"Kavuda/utils"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func (d AdaDeranaDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}
	//convert html string to doc for element selection
	doc, err := utils2.HTMLStringToDoc(resp)
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
			dateString := strings.Replace(nodeDoc.Find("span").First().Nodes[0].FirstChild.Data, "  ", " ", -1) // replacing double &nbsp; characters with space

			newsItems = append(newsItems, models.NewsItem{
				Title:   title,
				Link:    url,
				Date:    utils.ExtractPublishedDate("| January 2, 2006 3:04 pm", dateString),
				Snippet: snippet,
			})
		}
	}

	return newsItems, nil
}
