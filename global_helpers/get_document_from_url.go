package global_helpers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/request_handlers"
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
