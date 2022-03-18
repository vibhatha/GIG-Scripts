package helpers

import "GIG-SDK/models"

func ConvertLinksToEntities(entity models.Entity, queue chan string) []models.Entity {
	var (
		linkEntities []models.Entity
	)

	for _, link := range entity.Links {
		if link.GetTitle() != "" {
			if !visited[link.GetTitle()] {
				//log.Println("	passed link ->", link.Title)
				go func(title string) {
					queue <- title
					//log.Println("	queued link ->", link.Title)
				}(link.GetTitle())
			}
			//add link as an entity
			linkEntities = append(linkEntities, models.Entity{Title: link.GetTitle()})
		}
	}

	return linkEntities
}
