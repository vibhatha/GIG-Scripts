package main

import (
	"GIG-Scripts"
	"GIG-Scripts/orgchart/helpers"
	"flag"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
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

		entity, err = GIG_Scripts.GigClient.AddEntitiesAsLinks(entity, entities)
		if err != nil {
			panic(err)
		}

		//save to db
		entity, saveErr := GIG_Scripts.GigClient.CreateEntity(entity)
		libraries.ReportError(saveErr, ministry)
	}

}
