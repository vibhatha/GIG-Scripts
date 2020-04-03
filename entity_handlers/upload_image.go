package entity_handlers

import (
	"GIG-SDK/models"
	"GIG-Scripts"
	"GIG-SDK/request_handlers"
)

/**
Upload an image through API
 */
func UploadImage(payload models.Upload) error {

	if _, err := request_handlers.PostRequest(scripts.ApiUrl+"upload", payload); err != nil {
		return err
	}
	return nil
}
