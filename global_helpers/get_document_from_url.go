package global_helpers

import (
	GIG_Scripts "GIG-Scripts"
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/libraries"
)

func GetDocumentFromUrl(link string) (*goquery.Document, error) {
	//get the page
	resp, err := GIG_Scripts.GigClient.GetRequest(link)
	if err != nil {
		return nil, err
	}
	//convert html string to doc for element selection
	return libraries.HTMLStringToDoc(resp)
}
