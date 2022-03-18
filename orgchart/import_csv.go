package main

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"GIG-Scripts/orgchart/helpers"
	"flag"
	"log"
	"os"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("file path not specified")
		os.Exit(1)
	}

	dataArray, fileName, gazetteDate, err := helpers.LoadDataFromCsv(args[0])

	if err != nil {
		log.Fatal("error loading csv file:", err)
	}

	dataStructure, nameStructure := helpers.GenerateDataStructures(dataArray, fileName)

	for ministry, departments := range dataStructure {
		log.Println(ministry)

		entity := helpers.GenerateEntityFromDataRecord(fileName, ministry, departments, gazetteDate, nameStructure)

		var entities []models.Entity
		for _, departmentName := range departments {
			childEntity := helpers.CreateChildEntity(fileName, departmentName, gazetteDate, ministry)
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
