package helpers

import (
	"GIG-Scripts"
	"log"
)

func SaveEntity(title string, queue chan string) {
	if title != "" {
		entity := Enqueue(title, queue)
		if !entity.IsNil() {
			_, err := GIG_Scripts.GigClient.CreateEntity(entity)
			if err != nil {
				log.Println(err.Error(), title)
			}
		}
	}
}