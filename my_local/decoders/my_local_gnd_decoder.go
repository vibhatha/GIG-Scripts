package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

const GNDDataSource = "gig-data-master/geo/gnd/"

type MyLocalGNDDecoder struct {
	MyLocalDecoder
}

func (d MyLocalGNDDecoder) DecodeToEntity(record []string, source string) models.Entity {

	//0-id	1-gnd_id	2-gnd_num	3-name	4-province_id	5-district_id	6-dsd_id	7-centroid	8-population
	parentId := record[5]
	parentEntity, err := GIG_Scripts.GigClient.GetEntityByAttribute("attributes.location_id", parentId)
	if err != nil {
		log.Fatal("err fetching parent entity", parentId)
	}

	entity := *new(extended_models.Location).
		SetLocationId(record[6], source).
		SetName(record[3], source).
		SetCentroid(record[7], source).
		SetPopulation(record[8], source).
		SetParent(parentEntity.GetTitle(), source).
		SetGeoCoordinates(DSDDataSource + record[6] + ".json").
		AddCategory("Grama Niladhari Division").AddLink(models.Link{Title: parentEntity.GetTitle()})

	//update parent entity
	payload := models.UpdateEntity{
		SearchAttribute: "attributes.location_id",
		SearchValue:     *new(models.Value).SetValueString(parentId),
		Attribute:       "grama_niladhari_divisions",
		Value:           *new(models.Value).SetSource(source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.AppendToEntity(payload); err != nil {
		log.Fatal("error updating parent entity", err)
	}

	return entity
}
