package helpers

import (
	"GIG-Scripts"
	"GIG-Scripts/press-releases/constants"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"strings"
	"time"
)

var categories = []string{"Press Releases"}

func CreateEntityFromImage(img string) (models.Entity, []models.NERResult, time.Time) {
	imageUrl := GetImageUrl(img)
	releaseDate, err := GetTime(img)
	if err != nil {
		log.Println(err)
		panic("invalid filename")
	}

	textContent, err := GIG_Scripts.GigClient.GetRequest(GIG_Scripts.GigClient.OcrServerUrl + imageUrl)
	if err != nil {
		panic(err)
	}
	//NER extraction
	entityTitles, err := GIG_Scripts.GigClient.ExtractEntityNames(textContent)
	libraries.ReportError(err)
	title := strings.Replace(img, "/images/", "", -1)
	title = strings.Replace(title, "/", "_", -1)

	entity := models.Entity{
		Title: title,
	}
	entity.SetSource(constants.SourceString+img).
		SetSourceDate(releaseDate).
		SetSourceSignature("trusted").
		AddCategories(categories).
		SetAttribute("extracted_text", models.Value{
			ValueType:   "string",
			ValueString: textContent,
			Date:        releaseDate,
			Source:      constants.SourceString + img,
			UpdatedAt:   time.Now(),
		})

	return entity, entityTitles, releaseDate
}
