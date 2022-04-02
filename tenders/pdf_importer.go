package main

import (
	"GIG-Scripts"
	"flag"
	"github.com/lsflk/gig-sdk/libraries"
	"log"
	"os"
)

/**
config before running
 */

var pdfCategories = []string{"Gazette"}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]
	//parse pdf
	textContent := libraries.ParsePdf(filePath)
	entityTitles, err := GIG_Scripts.GigClient.ExtractEntityNames(textContent)
	libraries.ReportError(err)
	if err := GIG_Scripts.GigClient.CreateEntityFromText(textContent, "Gazette 2015", pdfCategories, entityTitles); err != nil {
		log.Println(err.Error(), filePath)
	} else {
		log.Println("pdf importing completed")
	}
}
