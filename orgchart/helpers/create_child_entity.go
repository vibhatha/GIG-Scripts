package helpers

import (
	"GIG-Scripts/orgchart/constants"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

var childCategories = []string{"Organization", "OrgChart-Level1"}

func CreateChildEntity(fileName string, departmentName string, gazetteDate time.Time, ministry string) models.Entity {
	childEntity := models.Entity{}
	childEntity.
		SetTitle(models.Value{ValueType: ValueType.String, ValueString: departmentName, Source: constants.SourceName + fileName, Date: gazetteDate}).
		SetSource(constants.SourceName+fileName).
		SetSourceDate(gazetteDate).
		AddCategories(childCategories).AddLink(*new(models.Link).SetTitle(ministry).AddDate(gazetteDate)).
		SetAttribute("parent", models.Value{
			ValueType:   "string",
			ValueString: ministry,
			Date:        gazetteDate,
			Source:      constants.SourceName + fileName,
			UpdatedAt:   time.Now(),
		})

	return childEntity
}
