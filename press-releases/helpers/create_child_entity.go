package helpers

import (
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"GIG-Scripts/press-releases/constants"
	"time"
)

func CreateChildEntity(mentionedEntity models.NERResult, entity models.Entity, releaseDate time.Time, img string)models.Entity{
	models.Entity{}.
		SetTitle(models.Value{ValueType: ValueType.String, ValueString: mentionedEntity.GetEntityName(), Source:constants.SourceString + img, Date: releaseDate}).
		SetSource(constants.SourceString + img).
		SetSourceDate(releaseDate).
		AddCategory(mentionedEntity.GetCategory()).AddLink(models.Link{}.SetTitle(entity.Title).AddDate(releaseDate)).
		SetAttribute("source", models.Value{
			ValueType:   "string",
			ValueString: constants.SourceName,
			Date:        releaseDate,
			Source:      constants.SourceString+ img,
			UpdatedAt:   time.Now(),
		})
}
