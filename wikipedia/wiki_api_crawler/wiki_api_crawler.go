// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"GIG-Scripts/wikipedia/wiki_api_crawler/helpers"
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	log.Println(args)
	if len(args) < 1 {
		log.Println("starting title not specified")
		os.Exit(1)
	}
	queue := make(chan string)
	go func() { queue <- args[0] }()

	for title := range queue {
		helpers.SaveEntity(title, queue)
	}
	log.Println("end")
}
