package main

import (
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"GIG-Scripts/tenders/etender/decoders"
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
)

var category = "Tenders"

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

			companyEntity := models.Entity{
			}.SetTitle(models.Value{
				ValueType:   "string",
				ValueString: tender.Company,
				Source:      tender.SourceName,
				Date:        tender.SourceDate,
			}).AddCategories([]string{"Organization", "Tenders",
			})

			locationEntity := models.Entity{
			}.SetTitle(models.Value{
				ValueType:   "string",
				ValueString: tender.Location,
				Source:      tender.SourceName,
				Date:        tender.SourceDate,
			}).AddCategory("Location")

			entity, _, addCompanyError := request_handlers.AddEntityAsAttribute(entity, "Company", companyEntity)
			if addCompanyError != nil {
				log.Println(addCompanyError)
			}
			entity, _, addLocationError := request_handlers.AddEntityAsAttribute(entity, "Location", locationEntity)
			if addLocationError != nil {
				log.Println(addLocationError)
			}

			savedEntity, saveErr := request_handlers.CreateEntity(entity)

			if saveErr != nil {
				log.Println(saveErr.Error(), entity)
			}
			log.Println(savedEntity.Title)
		}
	}
}
