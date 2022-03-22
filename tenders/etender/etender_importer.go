package main

import (
	"GIG-Scripts"
	"GIG-Scripts/global_helpers"
	"GIG-Scripts/tenders/etender/constants"
	"GIG-Scripts/tenders/etender/decoders"
	"GIG-Scripts/tenders/etender/helpers"
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
)

var category = constants.Tenders

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]

	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	ignoreHeaders := true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if ignoreHeaders {
			ignoreHeaders = false
		} else {
			tender := decoders.Decode(line)

			entity := decoders.MapToEntity(tender).AddCategory(category)
			companyEntity := helpers.CreateCompanyEntity(tender)
			locationEntity := helpers.CreateLocationEntity(tender)

			entity, _, addCompanyError := GIG_Scripts.GigClient.AddEntityAsAttribute(entity, constants.Company, companyEntity)
			global_helpers.ReportErrorWithoutCrash(addCompanyError)

			entity, _, addLocationError := GIG_Scripts.GigClient.AddEntityAsAttribute(entity, constants.Location, locationEntity)
			global_helpers.ReportErrorWithoutCrash(addLocationError)

			savedEntity, saveErr := GIG_Scripts.GigClient.CreateEntity(entity)
			global_helpers.ReportErrorWithoutCrash(saveErr)
			log.Println(savedEntity.Title)
		}
	}
}
