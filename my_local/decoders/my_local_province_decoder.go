package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

const ProvinceDataSource = "gig-data-master/geo/province/"

type MyLocalProvinceDecoder struct {
	MyLocalDecoder
}

func (d MyLocalProvinceDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(extended_models.Location).SetLocationId(record[1], source).
		SetName(record[2]+" Province", source).
		SetCentroid(record[3], source).
		SetPopulation(record[4], source).
		SetParent("Sri Lanka", source).
		SetGeoCoordinates(ProvinceDataSource + record[1] + ".json").
		AddCategory("Province").AddLink(models.Link{Title: "Sri Lanka"})

	//update parent entity
	payload := models.UpdateEntity{
		//Title:     "Sri Lanka",
		SearchAttribute: "attributes.location_id",
		SearchValue:     *new(models.Value).SetValueString("LK"),
		Attribute:       "provinces_test",
		Value:           *new(models.Value).SetSource(source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(payload); err != nil {
		log.Fatal("error updating parent entity", err)
	}

	return entity
}
