package entity_handlers

import (
	"GIG-Scripts/entity_handlers"
	"GIG-SDK/models"
	"fmt"
)

func (t *TestEntityHandlers) TestThatAddEntityAsAttributeWorks() {
	attributeEntity, err := entity_handlers.GetEntity("Sri Lanka")
	if err != nil {
		t.AssertNotFound()
	}
	entity := models.Entity{Title: "test entity"}
	entity, attributeEntity, _ = entity_handlers.AddEntityAsAttribute(entity, "testAttribute", attributeEntity)
	fmt.Println(attributeEntity.Id)
	fmt.Println(entity.Attributes[0].Values)
	t.AssertEqual(entity.Attributes[0].Values[0].ValueString, attributeEntity.Id.Hex())

}