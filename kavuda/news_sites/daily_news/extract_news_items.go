package daily_news

import (
	"GIG-SDK/libraries"
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func (d DailyNewsDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	var allNewsItems []models.NewsItem

	for _, newsSource := range newsSources {
		newsItems, err := extractNewItems(newsSource)
		if err != nil {
			log.Fatal("error loading new items ")
		} else {
			allNewsItems = append(allNewsItems, newsItems...)
		}
	}

	return allNewsItems, nil
}

func extractNewItems(newsSource models.NewsSource) ([]models.NewsItem, error) {
	doc, err := global_helpers.GetDocumentFromUrl(newsSource.Link)
	if err != nil {
		return nil, err
	}

	var newsLinks []string

	newsNodes := doc.Find(".field-content")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		extractedUrl, exist := nodeDoc.Find("a").First().Attr("href")

		if exist { // if url found
			title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
			if title != "img" { //is valid news link
				url := libraries.FixUrl(extractedUrl, newsSource.Link)

				if !libraries.StringInSlice(newsLinks, url) && !strings.Contains(url, "#comment") { // if the link is not already enlisted before
					newsLinks = append(newsLinks, url)

					extractDate := strings.Split(extractedUrl, "/")
					dateString := extractDate[1] + " " + extractDate[2] + " " + extractDate[3]
					newsItems = append(newsItems, models.NewsItem{
						Title:      title,
						Link:       url,
						Date:       helpers.ExtractPublishedDate("2006 01 02", dateString),
						Categories: newsSource.Categories,
					})
				}
			}
		}
	}

	return newsItems, nil
}
