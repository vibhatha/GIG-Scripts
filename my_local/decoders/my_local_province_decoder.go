package decoders

import (
	"GIG-Scripts/my_local/data_models"
	"github.com/lsflk/gig-sdk/models"
)

const GeoDataSource = "gig-data-master/geo/province/"

type MyLocalProvinceDecoder struct {
	MyLocalDecoder
}

func (d MyLocalProvinceDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(data_models.Province).SetLocationId(record[1], source).
		SetName(record[2]+" Province", source).
		SetCentroid(record[3], source).
		SetPopulation(record[4], source).
		SetParent("Sri Lanka", source).
		SetGeoCoordinates(GeoDataSource + record[1] + ".json").
		AddCategory("Province").AddLink(models.Link{Title: "Sri Lanka"})
	return entity
}
