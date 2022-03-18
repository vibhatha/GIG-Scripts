package helpers

import (
	"GIG-SDK/libraries"
	"encoding/csv"
	"log"
	"os"
	"time"
)

func LoadDataFromCsv(filePath string) ([][]string, string, time.Time, error) {
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

	return dataArray, fileName, gazetteDate, err
}
