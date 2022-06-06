package main

import (
	"GIG-Scripts/my_local/decoders"
	"GIG-Scripts/my_local/helpers"
)

const (
	DataPath = "gig-data-master/"
)

func main() {
	// open file
	countrySource := DataPath + "country.tsv"
	provinceSource := DataPath + "province.tsv"

	helpers.AddDecodedData(countrySource, decoders.MyLocalCountryDecoder{})
	helpers.AddDecodedData(provinceSource, decoders.MyLocalProvinceDecoder{})
}
