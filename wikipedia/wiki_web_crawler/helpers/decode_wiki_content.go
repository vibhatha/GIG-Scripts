package helpers

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/wikipedia/wiki_web_crawler/parsers"
	"github.com/lsflk/gig-sdk/libraries/clean_html"
	"github.com/lsflk/gig-sdk/models"
	"golang.org/x/net/html"
)

func DecodeWikiContent(uri string) (models.Entity, []models.Entity, []models.Upload, error) {
	var (
		wikiArticle extended_models.WikipediaArticle
		err         error
		body        *html.Node
	)

	wikiArticle.SetSource(uri).SetSourceSignature("trusted")

	doc, err := global_helpers.GetDocumentFromUrl(uri)
	if err != nil {
		return wikiArticle.Entity, []models.Entity{}, []models.Upload{}, err
	}

	wikiArticle.Title, body, err = parsers.ParseHTMLContent(doc)
	if err != nil {
		return wikiArticle.Entity, []models.Entity{}, []models.Upload{}, err
	}

	//clean html code by removing unwanted information
	htmlCleaner := clean_html.HtmlCleaner{Config: clean_html.Config{
		LineBreakers:   []string{"div", "caption"},
		IgnoreElements: []string{"noscript", "script", "style", "input"},
		IgnoreStrings:  []string{"[", "]", "edit", "Jump to search", "Jump to navigation"},
		IgnoreTitles:   []string{"(page does not exist)", ":"},
		IgnoreClasses:  []string{"box-Multiple_issues"},
	}}
	result, linkedEntities, imageList, defaultImageSource := htmlCleaner.CleanHTML(uri, body)
	wikiArticle.ImageURL = defaultImageSource
	wikiArticle.SetContent(result)

	return wikiArticle.Entity, linkedEntities, imageList, err
}
