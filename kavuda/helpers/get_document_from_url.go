package helpers

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	"github.com/PuerkitoBio/goquery"
)

func GetDocumentFromUrl(link string) (*goquery.Document, error) {
	//get the page
	resp, err := request_handlers.GetRequest(link)
	if err != nil {
		return nil, err
	}
	//convert html string to doc for element selection
	return libraries.HTMLStringToDoc(resp)
}
