package decoders

import (
	"GIG-Scripts/extended_models"
	"time"
)

func Decode(result []string) extended_models.Tender {
	sourceDate, _ := time.Parse("01/02/06", result[2])
	closingDate, _ := time.Parse("01/02/06", result[6])
	tender := extended_models.Tender{}
	tender.
		SetTenderTitle(result[0], result[7], sourceDate).
		SetCompany(result[1]).
		SetClosingDate(closingDate).
		SetTenderValue(result[9]).
		SetDescription(result[8]).
		SetLocation(result[5]).
		SetSource(result[7]).
		SetSourceDate(sourceDate).
		AddCategory(result[3]).
		AddCategory(result[4])

	return tender
}
