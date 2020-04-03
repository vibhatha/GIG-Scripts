package entity_handlers

import (
	"GIG-SDK/models"
)

/**
Add entity as an link to a given entity
 */
func AddEntityAsLink(entity models.Entity, linkEntity models.Entity) (models.Entity, models.Entity, error) {
	createdLinkEntity, linkEntityCreateError := CreateEntity(linkEntity)
	if linkEntityCreateError != nil {
		return entity, createdLinkEntity, linkEntityCreateError
	}
	entity = entity.AddLink(models.Link{}.SetTitle(linkEntity.GetTitle()).AddDate(entity.GetSourceDate()))
	return entity, createdLinkEntity, nil
}