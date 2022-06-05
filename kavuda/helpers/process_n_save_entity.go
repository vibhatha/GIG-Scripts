package helpers

import (
	"GIG-Scripts"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

func ProcessAndSaveEntity(entity models.Entity, textContent string) {
	//NER extraction
	log.Println("		Running NER on the text content...")
	entityTitles, err := GIG_Scripts.GigClient.ExtractEntityNames(textContent)
	if err != nil {
		log.Println("		NER failed:", err, entity.Title, entityTitles)
	} else {
		log.Println("		NER completed successfully.")
	}

	var entities []models.Entity

	for _, entityObject := range entityTitles {
		normalizedName := entityObject.EntityName
		if err == nil {
			titleValue := models.Value{
				ValueType:   ValueType.String,
				ValueString: normalizedName,
				Source:      entity.Source,
				Date:        entity.SourceDate}
			entities = append(entities, *new(models.Entity).SetTitle(titleValue).AddCategory(entityObject.Category))
		}
	}

	err = GIG_Scripts.GigClient.AddEntitiesAsLinks(&entity, entities)
	//save to db
	entity, saveErr := GIG_Scripts.GigClient.CreateEntity(entity)
	if saveErr != nil {
		log.Println(saveErr.Error(), entity.Title)
	} else {
		log.Println("		News Article Saved.", entity.Title)
	}
}
