// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG-Scripts"
	"GIG-Scripts/wikipedia/wiki_web_crawler/helpers"
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

	queue := make(chan string)

	// initialize queue from log or command arguments
	if err := helpers.LoadQueueFromLog(queue); err != nil {
		args := helpers.CheckArgs()
		log.Println("initializing queue from command arguments")
		go func(url string) { queue <- url }(args[0])
	}

	for uri := range queue {
		select {
		case <-exit:
			helpers.GracefulShutdown(queue, 0)
		default:
			entity, err := enqueue(uri, queue)
			if err != nil {
				log.Println("enqueue error:", err.Error(), uri)
				helpers.GracefulShutdown(queue, 1)
			}
			_, err = GIG_Scripts.GigClient.CreateEntity(entity)
			if err != nil {
				log.Println("create entity error:", err.Error(), uri)
				helpers.GracefulShutdown(queue, 1)
			}
			log.Println("entity added", entity.Title)
		}
	}
}

func enqueue(uri string, queue chan string) (models.Entity, error) {
	log.Println("fetching", uri)
	visited[uri] = true

	entity, linkedEntities, imageList, err := helpers.DecodeWikiContent(uri)
	if err != nil {
		return entity, err
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
	return entity, err
}
