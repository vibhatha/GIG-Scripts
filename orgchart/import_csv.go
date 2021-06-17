package main

import (
	"GIG-SDK"
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/libraries"
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

var dataStructure = make(map[string][]string)
var categories = []string{"OrgChart"}
var childCategories = []string{"Organization", "OrgChart-Level1"}
var nameStructure = make(map[string]string)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]
	fileName := libraries.ExtractFileName(filePath)
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvFile)
	r.Comma = ','
	dataArray, err := r.ReadAll()
	if err != nil {
		log.Println(err)
		panic("error reading csv")
	}
	log.Println(fileName)

	gazetteDate, err := time.Parse("gazette-2006-1-2.csv", fileName)
	if err != nil {
		log.Println(err)
		panic("invalid filename")
	}
	log.Println(gazetteDate)

	if err != nil {
		log.Fatalln("Error reading the csv file", err)
	}

	for _, record := range dataArray {
		record0, record1 := strings.TrimSpace(record[0]), strings.TrimSpace(record[1])

		if record0 == "Terminate" {
			terminateEntities(fileName, record1)
		} else {
			ministryName := record0
			childName := record1

			//if ministry name change is detected, create a change name request
			ministryNameArray := strings.Split(ministryName, "->")
			if len(ministryNameArray) == 2 {
				ministryName = strings.TrimSpace(ministryNameArray[1])
				nameStructure[ministryName] = strings.TrimSpace(ministryNameArray[0])
			}

			dataStructure[ministryName] = append(dataStructure[ministryName], childName)
		}
	}

	for ministry, departments := range dataStructure {
		log.Println(ministry)
		var filteredDepartments []string
		for _, department := range departments {
			log.Println("	", department)
			if department != "" {
				filteredDepartments = append(filteredDepartments, department)
			}
		}

		//decode to entity
		var entities []models.Entity
		jsonDepartments, err := json.Marshal(filteredDepartments)
		if err != nil {
			panic("Error converting to json,")
		}

		ministryName := ministry
		if oldName, ok := nameStructure[ministry]; ok {
			ministryName = oldName
		}

		entity := models.Entity{
			Title: ministryName,
		}.
			SetSource("Gazette " + fileName).
			SetSourceDate(gazetteDate).
			SetSourceSignature("trusted").
			AddCategories(categories).
			SetAttribute("organizations",
				models.Value{
					ValueType:   "json",
					ValueString: string(jsonDepartments),
					Source:      "Gazette " + fileName,
					Date:        gazetteDate,
					UpdatedAt:   time.Now(),
				})

		// detect entity name changes and include it in attributes
		if _, newTitleFound := nameStructure[ministry]; newTitleFound {
			entity = entity.SetAttribute("new_title",
				models.Value{
					ValueType:   "string",
					ValueString: ministry,
					Source:      "Gazette " + fileName,
					Date:        gazetteDate,
					UpdatedAt:   time.Now(),
				})
		}

		for _, departmentName := range departments {
			childEntity := models.Entity{}.
				SetTitle(models.Value{ValueType: ValueType.String, ValueString: departmentName, Source: "Gazette " + fileName, Date: gazetteDate}).
				SetSource("Gazette " + fileName).
				SetSourceDate(gazetteDate).
				AddCategories(childCategories).AddLink(models.Link{}.SetTitle(ministry).AddDate(entity.GetSourceDate())).
				SetAttribute("parent", models.Value{
					ValueType:   "string",
					ValueString: ministry,
					Date:        gazetteDate,
					Source:      "Gazette " + fileName,
					UpdatedAt:   time.Now(),
				})

			entities = append(entities, childEntity)
		}

		entity, err = request_handlers.AddEntitiesAsLinks(entity, entities)
		if err != nil {
			panic(err)
		}

		//save to db
		entity, saveErr := request_handlers.CreateEntity(entity)
		if saveErr != nil {
			log.Println(err.Error(), ministry)
		}
	}

}

func terminateEntities(fileName string, entityName string) {
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
