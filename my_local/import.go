package main

import (
	"GIG-Scripts/my_local/data_models"
	"GIG-Scripts/my_local/helpers"
)

const (
	DataPath = "gig-data-master/"
)

func main() {
	// open file
	filename := DataPath + "country.tsv"
	helpers.ImportFile(filename, data_models.MyLocalCountryDecoder{})
}
