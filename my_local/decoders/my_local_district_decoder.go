package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/my_local/data_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

const DistrictDataSource = "gig-data-master/geo/district/"

type MyLocalDistrictDecoder struct {
	MyLocalDecoder
}

func (d MyLocalDistrictDecoder) DecodeToEntity(record []string, source string) models.Entity {
	parentId := record[2]
	parentEntity, err := GIG_Scripts.GigClient.GetEntityByAttribute("attributes.location_id", parentId)
	if err != nil {
		log.Fatal("err fetching parent entity", parentId)
	}

	entity := *new(data_models.Province).SetLocationId(record[1], source).
		SetName(record[3]+" District", source).
		SetCentroid(record[4], source).
		SetPopulation(record[5], source).
		SetParent(parentEntity.GetTitle(), source).
		SetGeoCoordinates(DistrictDataSource + record[1] + ".json").
		AddCategory("District").AddLink(models.Link{Title: parentEntity.GetTitle()})

	//update parent entity
	payload := models.UpdateEntity{
		//Title:     "Sri Lanka",
		SearchAttribute: "attributes.location_id",
		SearchValue:     *new(models.Value).SetValueString(parentId),
		Attribute:       "districts",
		Value:           *new(models.Value).SetSource(source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(payload); err != nil {
		log.Fatal("error updating parent entity", err)
	}

	return entity
}
