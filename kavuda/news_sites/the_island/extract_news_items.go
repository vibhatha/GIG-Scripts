package the_island

import (
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/kavuda/helpers"
	"GIG-Scripts/kavuda/models"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/libraries"
	"golang.org/x/net/html"
)

func (d TheIslandDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	var allNewsItems []models.NewsItem

	for _, newsSource := range newsSources {
		doc, err := global_helpers.GetDocumentFromUrl(newsSource.Link)
		if err != nil {
			return nil, err
		}

		var newsLinks []string

		newsNodes := doc.Find(".mvp-blog-story-wrap")

		var newsItems []models.NewsItem
		for _, node := range newsNodes.Nodes {
			newsItem, url, err := generateNewsItem(node, newsSource)
			if !libraries.StringInSlice(newsLinks, url) && err == nil { // if the link is not already enlisted before
				newsLinks = append(newsLinks, url)
				newsItems = append(newsItems, newsItem)
			}
		}
		allNewsItems = append(allNewsItems, newsItems...)
	}

	return allNewsItems, nil
}

func generateNewsItem(node *html.Node, newsSource models.NewsSource) (models.NewsItem, string, error) {
	nodeDoc := goquery.NewDocumentFromNode(node)
	dateString, _ := nodeDoc.Find(".mvp-cd-date").First().Html()

	if dateString != "" {
		extractedUrl, _ := nodeDoc.Find("a").First().Attr("href")
		if extractedUrl != "/" {
			title := nodeDoc.Find("h2").First().Text()
			url := libraries.FixUrl(extractedUrl, newsSource.Link)

			newsItem := models.NewsItem{
				Title:      title,
				Link:       url,
				Date:       helpers.ExtractPublishedDate("January 2, 2006, 3:04 pm", dateString),
				Categories: newsSource.Categories,
			}
			return newsItem, url, nil

		}
	}

	return models.NewsItem{}, "", errors.New("news item not found")
}
