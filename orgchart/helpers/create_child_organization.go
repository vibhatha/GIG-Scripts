package helpers

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/orgchart/constants"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

var childCategories = []string{"Organization", "OrgChart-Level1"}

func CreateChildOrganization(fileName string, departmentName string, gazetteDate time.Time, ministry string) extended_models.Organization {
	childEntity := extended_models.Organization{}
	childEntity.
		SetParentOrganization(ministry, constants.SourceName+fileName, gazetteDate).
		SetTitle(models.Value{ValueType: ValueType.String, ValueString: departmentName, Source: constants.SourceName + fileName, Date: gazetteDate}).
		SetSource(constants.SourceName + fileName).
		SetSourceDate(gazetteDate).
		AddCategories(childCategories).AddLink(*new(models.Link).SetTitle(ministry).AddDate(gazetteDate))

	return childEntity
}
