package data_models

import "github.com/lsflk/gig-sdk/models"

type MyLocalProvinceDecoder struct {
	MyLocalDecoder
}

func (d MyLocalProvinceDecoder) DecodeToEntity(record []string, source string) models.Entity {
	entity := *new(Province).SetProvinceId(record[1], source).
		SetName(record[2], source).
		SetCentroid(record[3], source).
		SetPopulation(record[3], source)
	return entity.Entity
}
