package main

import (
	"GIG-Scripts"
	"GIG-Scripts/press-releases/helpers"
	"flag"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"os"
)

/**
input a web url containing pdf files
automatically download and create entities for all of pdf files
 */

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("starting url not specified")
		os.Exit(1)
	}
	uri := args[0]

	resp, err := GIG_Scripts.GigClient.GetRequest(uri)

	if err != nil {
		panic(err)
	}
	images := helpers.FindImages(resp)

	for _, img := range images {
		if helpers.ImageIsFound(img) {

			entity, entityTitles, releaseDate := helpers.CreateEntityFromImage(img)

			var entities []models.Entity
			for _, mentionedEntity := range entityTitles {
				childEntity := helpers.CreateChildEntity(mentionedEntity, entity, releaseDate, img)
				entities = append(entities, childEntity)
			}

			entity, err = GIG_Scripts.GigClient.AddEntitiesAsLinks(entity, entities)
			if err != nil {
				panic(err)
			}

			//save to db
			entity, saveErr := GIG_Scripts.GigClient.CreateEntity(entity)
			libraries.ReportError(saveErr, img)
		}
	}

	log.Println("image crawling completed")

}
