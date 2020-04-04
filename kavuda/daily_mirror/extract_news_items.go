package daily_mirror

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	"Kavuda/models"
	"Kavuda/utils"
	"github.com/PuerkitoBio/goquery"
)

func (d DailyMirrorDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}

	//convert html string to doc for element selection
	doc, err := libraries.HTMLStringToDoc(resp)
	if err != nil {
		return nil, err
	}

	var newsLinks []string

	newsNodes := doc.Find(".col-md-8")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)

		url, _ := nodeDoc.Find("a").First().Attr("href")
		if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
			newsLinks = append(newsLinks, url)

			dateString := nodeDoc.Find(".gtime").First().Nodes[0].FirstChild.Data
			title := nodeDoc.Find(".cat-hd-tx").First().Nodes[0].FirstChild.Data
			snippet := nodeDoc.Find("p").Last().Nodes[0].FirstChild.Data

			newsItems = append(newsItems, models.NewsItem{
				Title:   title,
				Snippet: snippet,
				Link:    url,
				Date:    utils.ExtractPublishedDate("02 Jan 2006 ", dateString),
			})
		}
	}

	return newsItems, nil
}
