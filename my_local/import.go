package main

import (
	"GIG-Scripts/my_local/decoders"
	"GIG-Scripts/my_local/helpers"
)

const (
	DataPath = "gig-data-master/"
)

func main() {
	// TODO: graceful shutdown
	// open file
	// country_id	province_id	district_id	dsd_id	gnd_id	ed_id	pd_id	lg_name		lg_id	moh_id
	//countrySource := DataPath + "country.tsv"
	//provinceSource := DataPath + "province.tsv"
	//districtSource := DataPath + "district.tsv"
	//dsdSource := DataPath + "dsd.tsv"
	//gndSource := DataPath + "gnd.tsv"
	edSource := DataPath + "ed.tsv"

	//helpers.AddDecodedData(countrySource, decoders.MyLocalCountryDecoder{})
	//helpers.AddDecodedData(provinceSource, decoders.MyLocalProvinceDecoder{})
	//helpers.AddDecodedData(districtSource, decoders.MyLocalDistrictDecoder{})
	//helpers.AddDecodedData(dsdSource, decoders.MyLocalDSDDecoder{})
	//helpers.AddDecodedData(gndSource, decoders.MyLocalGNDDecoder{})
	helpers.AddDecodedData(edSource, decoders.MyLocalEDDecoder{})
}
