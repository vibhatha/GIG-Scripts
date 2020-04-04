package main

import (
	"GIG-SDK/libraries"
	"GIG-SDK/request_handlers"
	"flag"
	"fmt"
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
		fmt.Println("file path not specified")
		os.Exit(1)
	}
	filePath := args[0]
	//parse pdf
	textContent := libraries.ParsePdf(filePath)
	entityTitles, err := request_handlers.ExtractEntityNames(textContent)
	if err != nil {
		fmt.Println(err)
	}
	if err := request_handlers.CreateEntityFromText(textContent, "Gazette 2015", pdfCategories, entityTitles); err != nil {
		fmt.Println(err.Error(), filePath)
	}

	fmt.Println("pdf importing completed")

}
