package main

import (
	"GIG-Scripts/press-releases/helpers"
	"flag"
	"github.com/lsflk/gig-sdk/models"
	"github.com/lsflk/gig-sdk/request_handlers"
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

	resp, err := request_handlers.GetRequest(uri)

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

			entity, err = request_handlers.AddEntitiesAsLinks(entity, entities)
			if err != nil {
				panic(err)
			}

			//save to db
			entity, saveErr := request_handlers.CreateEntity(entity)
			if saveErr != nil {
				log.Println(err.Error(), img)
			}
		}
	}

	log.Println("image crawling completed")

}
