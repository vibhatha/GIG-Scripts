package decoders

import (
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/models"
)

type MyLocalCountryDecoder struct {
	MyLocalDecoder
}

func (d MyLocalCountryDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(extended_models.Location).
		SetLocationId(record[1], source).
		SetName(record[2], source).
		SetPopulation(record[3], source).
		AddCategory("Country")
	return entity
}
