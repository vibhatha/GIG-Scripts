package helpers

import (
	"GIG-Scripts/orgchart/constants"
	"encoding/json"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"time"
)

var categories = []string{constants.OrgChartCategory}

func GenerateEntityFromDataRecord(fileName string, ministry string, departments []string, gazetteDate time.Time, nameStructure map[string]string) models.Entity {
	var filteredDepartments []string
	for _, department := range departments {
		log.Println("	", department)
		if department != "" {
			filteredDepartments = append(filteredDepartments, department)
		}
	}

	//decode to entity
	jsonDepartments, err := json.Marshal(filteredDepartments)
	if err != nil {
		panic("Error converting to json,")
	}

	ministryName := ministry
	if oldName, ok := nameStructure[ministry]; ok {
		ministryName = oldName
	}

	entity := models.Entity{Title: ministryName}
	entity.SetSource(constants.SourceName+fileName).
		SetSourceDate(gazetteDate).
		SetSourceSignature("trusted").
		AddCategories(categories).
		SetAttribute("organizations",
			models.Value{
				ValueType:   "json",
				ValueString: string(jsonDepartments),
				Source:      constants.SourceName + fileName,
				Date:        gazetteDate,
				UpdatedAt:   time.Now(),
			})

	// detect entity name changes and include it in attributes
	if _, newTitleFound := nameStructure[ministry]; newTitleFound {
		entity.SetAttribute("new_title",
			models.Value{
				ValueType:   "string",
				ValueString: ministry,
				Source:      constants.SourceName + fileName,
				Date:        gazetteDate,
				UpdatedAt:   time.Now(),
			})
	}
	return entity
}
