// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG-Scripts"
	"GIG-Scripts/wikipedia/wiki_web_crawler/helpers"
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var visited = make(map[string]bool)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	args := helpers.CheckArgs()

	queue := make(chan string)
	go func() { queue <- args[0] }()

	for uri := range queue {
		select {
		case <-exit:
			log.Println("termination detected. Saving current progress...")
			//log.Println("exiting now.1")
			//os.Exit(0)
			for {
				select {
				case url := <-queue:
					log.Println(url)
				default:
					log.Println("exiting now.")
					os.Exit(0)
				}
			}
		default:
			entity, err := enqueue(uri, queue)
			if err != nil {
				log.Println("enqueue error:", err.Error(), uri)
			}
			log.Println(entity.ImageURL)
			_, err = GIG_Scripts.GigClient.CreateEntity(entity)
			log.Println("entity added", entity.Title)
			libraries.ReportError(err, uri)
		}
	}
}

func enqueue(uri string, queue chan string) (models.Entity, error) {
	log.Println("fetching", uri)
	visited[uri] = true

	entity, linkedEntities, imageList, err := helpers.DecodeWikiContent(uri)
	if err != nil {
		log.Fatal("error decoding response to entity")
	}
	helpers.UploadImages(imageList)

	// queue new links for crawling
	for _, linkedEntity := range linkedEntities {
		if !visited[linkedEntity.GetSource()] {
			go func(url string) {
				queue <- url
			}(linkedEntity.GetSource())
		}
		newLink := models.Link{}
		newLink.SetTitle(linkedEntity.GetTitle()).AddDate(entity.GetSourceDate())
		entity.AddLink(newLink)
	}

	// save linkedEntities (create empty if not exist)
	err = GIG_Scripts.GigClient.AddEntitiesAsLinks(&entity, linkedEntities)
	if err != nil {
		log.Fatal(err)
	}

	return entity, nil
}
