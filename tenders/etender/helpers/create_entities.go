package helpers

import (
	"GIG-Scripts/extended_models"
	"GIG-Scripts/tenders/etender/constants"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

func CreateCompanyEntity(tender extended_models.Tender) models.Entity {
	entity := models.Entity{}
	entity.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: tender.GetCompany(),
		Source:      tender.Source,
		Date:        tender.SourceDate,
	}).AddCategories([]string{constants.Organization, constants.Tenders})
	return entity
}

func CreateLocationEntity(tender extended_models.Tender) models.Entity {
	entity := models.Entity{}
	entity.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: tender.GetLocation(),
		Source:      tender.Source,
		Date:        tender.SourceDate,
	}).AddCategory(constants.Location)

	return entity
}
