package helpers

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"log"
)

func ProcessAndSaveEntity(entity models.Entity, textContent string) {
	//NER extraction
	log.Println("		Running NER on the text content...")
	entityTitles, err := request_handlers.ExtractEntityNames(textContent)
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

	entity, err = request_handlers.AddEntitiesAsLinks(entity, entities)
	//save to db
	entity, saveErr := request_handlers.CreateEntity(entity)
	if saveErr != nil {
		log.Println(saveErr.Error(), entity.Title)
	}
	log.Println("		News Article Saved.", entity.Title)
}
