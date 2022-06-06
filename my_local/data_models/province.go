package data_models

import (
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

type Province struct {
	extended_models.Location
}

func (p *Province) SetProvinceId(countryId string, source string) *Province {
	p.SetAttribute("province_id", models.Value{
		ValueType:   ValueType.String,
		ValueString: countryId,
		Source:      source,
	})
	return p
}
