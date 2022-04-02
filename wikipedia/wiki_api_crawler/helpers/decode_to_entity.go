package helpers

import (
	"GIG-Scripts/wikipedia/wiki_api_crawler/decoders"
	"GIG-Scripts/wikipedia/wiki_api_crawler/requests"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"sync"
)

func DecodeToEntity(title string) models.Entity {
	entity := models.Entity{}
	var requestWorkGroup sync.WaitGroup
	for _, propType := range requests.PropTypes() {

		requestWorkGroup.Add(1)
		go func(prop string) {
			defer requestWorkGroup.Done()
			result, err := requests.GetContent(prop, title)
			libraries.ReportError(err)
			if err == nil {
				decoders.Decode(result, &entity)
			}
		}(propType)
	}
	requestWorkGroup.Wait()

	return entity
}
