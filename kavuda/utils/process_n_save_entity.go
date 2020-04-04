package utils

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"fmt"
)

func ProcessAndSaveEntity(entity models.Entity, textContent string) {
	//NER extraction
	fmt.Println("		Running NER on the text content...")
	entityTitles, err := request_handlers.ExtractEntityNames(textContent)
	if err != nil {
		fmt.Println(err, entity.Title)
		fmt.Println(entityTitles)
	}
	fmt.Println("		NER completed successfully.")

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
		fmt.Println(saveErr.Error(), entity.Title)
	}
	fmt.Println("		News Article Saved.", entity.Title)
}
