package decoders

import (
	"GIG-Scripts/my_local/data_models"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalCountryDecoder struct {
	MyLocalDecoder
}

func (d MyLocalCountryDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(data_models.Country).
		SetLocationId(record[1], source).
		SetName(record[2], source).
		SetPopulation(record[3], source).
		AddCategory("Country")
	return entity
}
