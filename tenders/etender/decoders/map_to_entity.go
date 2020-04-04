package decoders

import (
	"GIG-SDK/models"
	"GIG-Scripts/tenders/etender/model"
)

func MapToEntity(tender model.ETender) models.Entity {
	return models.Entity{
	}.
		AddCategory(tender.Category).
		AddCategory(tender.Subcategory).
		SetTitle(models.Value{
			ValueType:   "string",
			ValueString: tender.Title + " - " + tender.Location,
			Date:        tender.SourceDate,
			Source:      tender.SourceName,
		}).
		SetAttribute("Source Date", models.Value{
			ValueType:   "date",
			ValueString: tender.SourceDate.String(),
		}).
		SetAttribute("Category", models.Value{
			ValueType:   "string",
			ValueString: tender.Category,
		}).
		SetAttribute("Subcategory", models.Value{
			ValueType:   "string",
			ValueString: tender.Subcategory,
		}).
		SetAttribute("Closing Date", models.Value{
			ValueType:   "date",
			ValueString: tender.ClosingDate.String(),
		}).
		SetAttribute("Source Name", models.Value{
			ValueType:   "string",
			ValueString: tender.SourceName,
		}).
		SetAttribute("Description", models.Value{
			ValueType:   "string",
			ValueString: tender.Description,
		}).
		SetAttribute("Value", models.Value{
			ValueType:   "string",
			ValueString: tender.Value,
		})
}
