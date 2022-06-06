package data_models

import "github.com/lsflk/gig-sdk/models"

type MyLocalCountryDecoder struct {
	MyLocalDecoder
}

func (d MyLocalCountryDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(Country).
		SetCountryId(record[1], source).
		SetName(record[2], source).
		SetPopulation(record[3], source)
	return entity.Entity
}
