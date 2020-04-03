package entity_handlers

import (
	"GIG-Scripts"
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"encoding/json"
)

/**
GetEntity
 */
func GetEntity(title string) (models.Entity, error) {
	var entity models.Entity
	resp, err := request_handlers.GetRequest(scripts.ApiUrl + "get/" + title)
	if err != nil {
		return entity, err
	}
	json.Unmarshal([]byte(resp), &entity)
	return entity, err
}
