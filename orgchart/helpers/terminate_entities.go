package helpers

import (
	"GIG-SDK"
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"log"
	"time"
)

func TerminateEntities(fileName string, entityName string) {
	/**
	get all orgchart entities
	for each entity add lifeStatus attribute valued terminated by given date
	append 'terminated by date' to all entities
	save entities
	  */

	terminationDate, _ := time.Parse("gazette-2006-1-2.csv", fileName)
	entity := models.Entity{}.SetSource("Gazette " + fileName).SetSourceDate(terminationDate).AddCategory("OrgChart")

	if entityName != "All" {
		entity.Title = entityName
	}
	resp, err := request_handlers.PostRequest(config.ApiUrl+"terminate", entity)
	if err != nil {
		log.Println("entity termination error:", err)
	}
	log.Println("Entity Termination:", resp)
}
