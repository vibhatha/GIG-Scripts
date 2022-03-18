package helpers

import (
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"GIG-Scripts/tenders/etender/constants"
	"GIG-Scripts/tenders/etender/model"
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
