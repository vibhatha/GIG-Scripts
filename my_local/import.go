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
	//countrySource := DataPath + "country.tsv"
	//provinceSource := DataPath + "province.tsv"
	districtSource := DataPath + "district.tsv"

	//helpers.AddDecodedData(countrySource, decoders.MyLocalCountryDecoder{})
	//helpers.AddDecodedData(provinceSource, decoders.MyLocalProvinceDecoder{})
	helpers.AddDecodedData(districtSource, decoders.MyLocalDistrictDecoder{})
}
