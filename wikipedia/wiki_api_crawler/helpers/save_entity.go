package helpers

import (
	"GIG-SDK/request_handlers"
	"log"
)

func SaveEntity(title string, queue chan string) {
	if title != "" {
		entity := Enqueue(title, queue)
		if !entity.IsNil() {
			_, err := request_handlers.CreateEntity(entity)
			if err != nil {
				log.Println(err.Error(), title)
			}
		}
	}
}