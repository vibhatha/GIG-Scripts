package helpers

import (
	"GIG-Scripts/tenders/etender/constants"
	"GIG-Scripts/tenders/etender/model"
	"github.com/lsflk/gig-sdk/enums/ValueType"
	"github.com/lsflk/gig-sdk/models"
)

func CreateCompanyEntity(tender model.ETender) models.Entity {
	return models.Entity{
	}.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: tender.Company,
		Source:      tender.SourceName,
		Date:        tender.SourceDate,
	}).AddCategories([]string{constants.Organization, constants.Tenders,
	})
}

func CreateLocationEntity(tender model.ETender)models.Entity{
	return models.Entity{
	}.SetTitle(models.Value{
		ValueType:   ValueType.String,
		ValueString: tender.Location,
		Source:      tender.SourceName,
		Date:        tender.SourceDate,
	}).AddCategory(constants.Location)
}
