package the_island

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	"GIG-Scripts/kavuda/models"
	"GIG-Scripts/kavuda/utils"
	"github.com/PuerkitoBio/goquery"
)

func (d TheIslandDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	var allNewsItems []models.NewsItem

	for _, newsSource := range newsSources {
		//get the page
		resp, err := request_handlers.GetRequest(newsSource.Link)
		if err != nil {
			return nil, err
		}
		//convert html string to doc for element selection
		doc, err := libraries.HTMLStringToDoc(resp)
		if err != nil {
			return nil, err
		}

		var newsLinks []string

		newsNodes := doc.Find(".col")

		var newsItems []models.NewsItem
		for _, node := range newsNodes.Nodes {
			nodeDoc := goquery.NewDocumentFromNode(node)
			dateString, _ := nodeDoc.Find(".article_date").First().Html()

			if dateString != "" {
				extractedUrl, _ := nodeDoc.Find("a").First().Attr("href")
				if extractedUrl != "/" {
					title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
					url := libraries.FixUrl(extractedUrl, newsSource.Link)

					if !libraries.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
						newsLinks = append(newsLinks, url)

						newsItems = append(newsItems, models.NewsItem{
							Title: title,
							Link:  url,
							Date:  utils.ExtractPublishedDate("January 2, 2006, 3:04 pm", dateString),
							Categories: newsSource.Categories,
						})
					}
				}
			}
		}
		allNewsItems = append(allNewsItems, newsItems...)
	}

	return allNewsItems, nil
}
