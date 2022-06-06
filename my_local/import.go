package main

import (
	GIG_Scripts "GIG-Scripts"
	"GIG-Scripts/my_local/data_models"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	DATA_PATH = "gig-data-master/"
)

func main() {
	// open file
	filename := DATA_PATH + "country.tsv"
	source := "MyLocal - " + filename
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read tsv values using csv.Reader
	tsvReader := csv.NewReader(f)
	tsvReader.Comma = '\t'
	headers, err := tsvReader.Read()
	if err != nil {
		log.Panicf("headers not found in %s", filename)
	}
	log.Println("header found:", headers)

	for {
		rec, err := tsvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		fmt.Printf("%+v\n", rec)
		country := *new(data_models.Country).SetCountryId(rec[1], source).SetName(rec[2], source).SetPopulation(rec[3], source)
		//save to db
		_, saveErr := GIG_Scripts.GigClient.CreateEntity(country.Entity)
		if saveErr != nil {
			log.Fatal(saveErr)
		}
	}
}
