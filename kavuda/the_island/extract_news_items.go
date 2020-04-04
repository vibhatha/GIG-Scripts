package the_island

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	utils2 "GIG-Scripts/wikipedia/utils"
	"Kavuda/models"
	"Kavuda/utils"
	"github.com/PuerkitoBio/goquery"
)

func (d TheIslandDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}
	resp1, err := request_handlers.GetRequest(newsSiteUrl1)
	if err != nil {
		return nil, err
	}
	//convert html string to doc for element selection
	doc, err := utils2.HTMLStringToDoc(resp)
	if err != nil {
		return nil, err
	}
	doc1, err := utils2.HTMLStringToDoc(resp1)
	if err != nil {
		return nil, err
	}

	var newsLinks []string

	newsNodes := doc.Find(".col")
	newsNodes1 := doc1.Find(".col")

	newsNodesNodes := append(newsNodes.Nodes, newsNodes1.Nodes...)
	var newsItems []models.NewsItem
	for _, node := range newsNodesNodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		dateString, _ := nodeDoc.Find(".article_date").First().Html()

		if dateString != "" {
			extractedUrl, _ := nodeDoc.Find("a").First().Attr("href")
			if extractedUrl != "/" {
				title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
				url := libraries.FixUrl(extractedUrl, newsSiteUrl)

				if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
					newsLinks = append(newsLinks, url)

					newsItems = append(newsItems, models.NewsItem{
						Title: title,
						Link:  url,
						Date:  utils.ExtractPublishedDate("January 2, 2006, 3:04 pm", dateString),
					})
				}
			}
		}
	}

	return newsItems, nil
}
