package helpers

import (
	"GIG-Scripts/press-releases/constants"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

func CreateChildEntity(mentionedEntity models.NERResult, entity models.Entity, releaseDate time.Time, img string) models.Entity {
	childEntity := models.Entity{}
	childEntity.SetTitle(models.Value{ValueType: ValueType.String, ValueString: mentionedEntity.GetEntityName(), Source: constants.SourceString + img, Date: releaseDate}).
		SetSource(constants.SourceString+img).
		SetSourceDate(releaseDate).
		AddCategory(mentionedEntity.GetCategory()).AddLink(*new(models.Link).SetTitle(entity.Title).AddDate(releaseDate)).
		SetAttribute("source", models.Value{
			ValueType:   "string",
			ValueString: constants.SourceName,
			Date:        releaseDate,
			Source:      constants.SourceString + img,
			UpdatedAt:   time.Now(),
		})

	return childEntity
}
