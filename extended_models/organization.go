package extended_models

import (
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
	"time"
)

type Organization struct {
	models.Entity
}

func (o *Organization) SetOrganizationTitle(title string, source string, sourceDate time.Time) *Organization {
	o.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: title,
		Source:      source,
		Date:        sourceDate,
	})
	return o
}

func (o *Organization) SetParentOrganization(organization string, source string, sourceDate time.Time) *Organization {
	o.SetAttribute("parent", models.Value{
		ValueType:   ValueType.String,
		ValueString: organization,
		Date:        sourceDate,
		Source:      source,
		UpdatedAt:   time.Now(),
	})
	return o
}

func (o *Organization) SetChildOrganizations(organizationsArrayString string, source string, sourceDate time.Time) *Organization {
	o.SetAttribute("organizations",
		models.Value{
			ValueType:   ValueType.JSON,
			ValueString: organizationsArrayString,
			Source:      source,
			Date:        sourceDate,
			UpdatedAt:   time.Now(),
		})
	return o
}
