package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

const DSDDataSource = "gig-data-master/geo/dsd/"

type MyLocalDSDDecoder struct {
	MyLocalDecoder
}

func (d MyLocalDSDDecoder) DecodeToEntity(record []string, source string) models.Entity {
	parentId := record[4]
	parentEntity, err := GIG_Scripts.GigClient.GetEntityByAttribute("attributes.location_id", parentId)
	if err != nil {
		log.Fatal("err fetching parent entity", parentId)
	}

	entity := *new(extended_models.Location).
		SetLocationId(record[1], source).
		SetName(record[2], source).
		SetCentroid(record[5], source).
		SetPopulation(record[6], source).
		SetParent(parentEntity.GetTitle(), source).
		SetGeoCoordinates(DSDDataSource + record[1] + ".json").
		AddCategory("Divisional Secretariats Division").AddLink(models.Link{Title: parentEntity.GetTitle()})

	//update parent entity
	payload := models.UpdateEntity{
		//Title:     "Sri Lanka",
		SearchAttribute: "attributes.location_id",
		SearchValue:     *new(models.Value).SetValueString(parentId),
		Attribute:       "divisional_secretariats_divisions",
		Value:           *new(models.Value).SetSource(source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(payload); err != nil {
		log.Fatal("error updating parent entity", err)
	}

	return entity
}
