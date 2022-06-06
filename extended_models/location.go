package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

type Location struct {
	models.Entity
}

func (l *Location) SetName(name string, source string) *Location {
	l.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: name,
		Source:      source,
	})
	return l
}

func (l *Location) SetCentroid(centroid string, source string) *Location {
	l.SetTitle(models.Value{
		ValueType:   ValueType.JSON,
		ValueString: centroid,
		Source:      source,
	})
	return l
}

func (l *Location) SetPopulation(population string, source string) *Location {
	l.SetAttribute("population", models.Value{
		ValueType:   ValueType.String,
		ValueString: population,
		Source:      source,
	})
	return l
}
