package global_helpers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/client"
	"github.com/lsflk/gig-sdk/libraries"
)

func GetDocumentFromUrl(link string) (*goquery.Document, error) {
	//get the page
	resp, err := client.GigClient{}.GetRequest(link)
	if err != nil {
		return nil, err
	}
	//convert html string to doc for element selection
	return libraries.HTMLStringToDoc(resp)
}
