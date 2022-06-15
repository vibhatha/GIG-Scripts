package decoders

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/my_local/data_models"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

const GeoDataSource = "gig-data-master/geo/province/"

type MyLocalProvinceDecoder struct {
	MyLocalDecoder
}

type Payload struct {
	Title     string       `json:"title"`
	Attribute string       `json:"attribute"`
	Value     models.Value `json:"value"`
}

func (d MyLocalProvinceDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(data_models.Province).SetLocationId(record[1], source).
		SetName(record[2]+" Province", source).
		SetCentroid(record[3], source).
		SetPopulation(record[4], source).
		SetParent("Sri Lanka", source).
		SetGeoCoordinates(GeoDataSource + record[1] + ".json").
		AddCategory("Province").AddLink(models.Link{Title: "Sri Lanka"})

	//update parent entity
	payload := Payload{
		Title:     "Sri Lanka",
		Attribute: "provinces",
		Value:     *new(models.Value).SetSource(source).SetValueString(entity.GetTitle()),
	}
	if _, err := GIG_Scripts.GigClient.PostRequest(GIG_Scripts.GigClient.ApiUrl+"append", payload); err != nil {
		log.Fatal("error updating parent entity", err)
	}

	return entity
}
