package helpers

import (
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"GIG-Scripts/orgchart/constants"
	"time"
)

var childCategories = []string{"Organization", "OrgChart-Level1"}

func CreateChildEntity(fileName string, departmentName string, gazetteDate time.Time, ministry string) models.Entity {
	return models.Entity{}.
		SetTitle(models.Value{ValueType: ValueType.String, ValueString: departmentName, Source: constants.SourceName + fileName, Date: gazetteDate}).
		SetSource(constants.SourceName + fileName).
		SetSourceDate(gazetteDate).
		AddCategories(childCategories).AddLink(models.Link{}.SetTitle(ministry).AddDate(gazetteDate)).
		SetAttribute("parent", models.Value{
			ValueType:   "string",
			ValueString: ministry,
			Date:        gazetteDate,
			Source:      constants.SourceName + fileName,
			UpdatedAt:   time.Now(),
		})
}
