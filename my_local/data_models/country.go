package data_models

import (
	"GIG-Scripts/extended_models"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

type Country struct {
	extended_models.Location
}

func (c *Country) SetCountryId(countryId string, source string) *Country {
	c.SetAttribute("country_id", models.Value{
		ValueType:   ValueType.String,
		ValueString: countryId,
		Source:      source,
	})
	return c
}

func (c *Country) SetName(name string, source string) *Country {
	c.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: name,
		Source:      source,
	})
	return c
}

func (c *Country) SetPopulation(population string, source string) *Country {
	c.SetAttribute("population", models.Value{
		ValueType:   ValueType.String,
		ValueString: population,
		Source:      source,
	})
	return c
}
