package helpers

import (
	"GIG-SDK"
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"GIG-Scripts/press-releases/constants"
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

	textContent, err := request_handlers.GetRequest(config.OCRServer + imageUrl)
	if err != nil {
		panic(err)
	}
	//NER extraction
	entityTitles, err := request_handlers.ExtractEntityNames(textContent)
	if err != nil {
		log.Println(err)
	}
	title := strings.Replace(img, "/images/", "", -1)
	title = strings.Replace(title, "/", "_", -1)

	entity := models.Entity{
		Title: title,
	}.
		SetSource(constants.SourceString + img).
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
