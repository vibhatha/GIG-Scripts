package helpers

import (
	GIG_Scripts "GIG-Scripts"
	"github.com/lsflk/gig-sdk/models"
	"log"
)

func UploadImages(imageList []models.Upload) {
	for _, image := range imageList {
		go func(payload models.Upload) {
			uploadErr := GIG_Scripts.GigClient.UploadFile(payload)
			if uploadErr != nil {
				log.Println("Error uploading file:", payload.Title, uploadErr)
			}
		}(image)
	}
}
