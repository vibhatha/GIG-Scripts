package decoders

import (
	"github.com/lsflk/gig-sdk/models"
	"log"
	"strings"
)

func Decode(result map[string]interface{}, entity *models.Entity) {
	query := result["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})

	for _, page := range pages {

		pageObj := page.(map[string]interface{})

		if pageObj["extract"] != nil {
			log.Println("	decoding content...")

			entity.Title = pageObj["title"].(string)
			tempEntity := entity.SetAttribute("", models.Value{}.
				SetType("wikiText").
				SetValueString(pageObj["extract"].(string)))
			entity.Attributes = tempEntity.Attributes
		}

		if pageObj["links"] != nil {
			log.Println("	decoding links...")
			links := pageObj["links"].([]interface{})

			for _, link := range links {
				linkObj := link.(map[string]interface{})
				entity.Links = append(entity.Links, models.Link{}.SetTitle(linkObj["title"].(string)))
			}
		}

		if pageObj["categories"] != nil {
			log.Println("	decoding categories...")
			categories := pageObj["categories"].([]interface{})

			for _, category := range categories {
				categoryObj := category.(map[string]interface{})
				categoryString := strings.Replace(categoryObj["title"].(string), "Category:", "", -1)
				entity.Categories = append(entity.Categories, categoryString)
			}
		}

	}

}
