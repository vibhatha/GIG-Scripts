package create_entity

import (
	"GIG-SDK/models"
	"GIG-Scripts/crawlers/utils"
	"GIG-Scripts/entity_handlers"
)

func CreateEntityFromText(textContent string, title string, categories []string, entityTitles []utils.NERResult) error {
	//decode to entity
	var entities []models.Entity
	entity := models.Entity{}.
		SetTitle(models.Value{}.SetType("string").SetValueString("gazette")).
		SetAttribute("", models.Value{}.
			SetType("string").
			SetValueString(textContent)).
		AddCategories(categories)

	for _, entityObject := range entityTitles {
		//normalizedName, err := utils.NormalizeName(entityObject.EntityName)
		//if err == nil {
		entities = append(entities, models.Entity{Title: entityObject.EntityName}.AddCategory(entityObject.Category))
		//}
	}

	entity, err := entity_handlers.AddEntitiesAsLinks(entity, entities)
	if err != nil {
		panic(err)
	}

	//save to db
	entity, saveErr := entity_handlers.CreateEntity(entity)

	return saveErr
}
