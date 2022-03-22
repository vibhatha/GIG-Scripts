package helpers

import (
	"github.com/lsflk/gig-sdk/models"
	"github.com/lsflk/gig-sdk/request_handlers"
	"log"
)

var visited = make(map[string]bool)

func Enqueue(title string, queue chan string) models.Entity {
	log.Println("fetching", title)
	visited[title] = true

	entity := DecodeToEntity(title)
	if !entity.IsNil() {

		linkEntities := ConvertLinksToEntities(entity, queue)
		var err error
		entity, err = request_handlers.AddEntitiesAsLinks(entity, linkEntities)

		if err != nil {
			log.Println("error creating links:", err)
		}
	}
	return entity
}
