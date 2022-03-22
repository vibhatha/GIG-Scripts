package helpers

import (
	"GIG-Scripts"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

func ProcessAndSaveEntity(entity models.Entity, textContent string) {
	//NER extraction
	log.Println("		Running NER on the text content...")
	entityTitles, err := GIG_Scripts.GigClient.ExtractEntityNames(textContent)
	if err != nil {
		log.Println(err, entity.Title)
		log.Println(entityTitles)
	}
	log.Println("		NER completed successfully.")

	var entities []models.Entity

	for _, entityObject := range entityTitles {
		normalizedName := entityObject.EntityName
		if err == nil {
			entities = append(entities, models.Entity{Title: normalizedName}.AddCategory(entityObject.Category))
		}
	}

	entity, err = GIG_Scripts.GigClient.AddEntitiesAsLinks(entity, entities)
	//save to db
	entity, saveErr := GIG_Scripts.GigClient.CreateEntity(entity)
	if saveErr != nil {
		log.Println(saveErr.Error(), entity.Title)
	}
	log.Println("		News Article Saved.", entity.Title)
}
