package helpers

import (
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"time"
)

var childCategories = []string{"Organization", "OrgChart-Level1"}

func CreateChildEntity(fileName string, departmentName string, gazetteDate time.Time, ministry string) models.Entity {
	return models.Entity{}.
		SetTitle(models.Value{ValueType: ValueType.String, ValueString: departmentName, Source: "Gazette " + fileName, Date: gazetteDate}).
		SetSource("Gazette " + fileName).
		SetSourceDate(gazetteDate).
		AddCategories(childCategories).AddLink(models.Link{}.SetTitle(ministry).AddDate(gazetteDate)).
		SetAttribute("parent", models.Value{
			ValueType:   "string",
			ValueString: ministry,
			Date:        gazetteDate,
			Source:      "Gazette " + fileName,
			UpdatedAt:   time.Now(),
		})
}
