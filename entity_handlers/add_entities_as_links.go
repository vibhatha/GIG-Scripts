package entity_handlers

import "GIG-SDK/models"

/**
Add list of related entities to a given entity
 */
func AddEntitiesAsLinks(entity models.Entity, linkEntities []models.Entity) (models.Entity, error) {
	createdLinkEntities, linkEntityCreateError := CreateEntities(linkEntities)
	if linkEntityCreateError != nil {
		return entity, linkEntityCreateError
	}
	for _, linkEntity := range createdLinkEntities {
		entity = entity.AddLink(models.Link{}.SetTitle(linkEntity.GetTitle()).AddDate(entity.GetSourceDate()))
	}
	return entity, nil
}
