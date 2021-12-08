// https://jdanger.com/build-a-web-crawler-in-go.html
package main

import (
	"flag"
	"github.com/joncrlsn/dque"
	"log"
	"os"
)

var visited = make(map[string]bool)

// Item is what we'll be storing in the queue.  It can be any struct
// as long as the fields you want stored are public.
type Item struct {
	Name string
	Id   int
}

// ItemBuilder creates a new item and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func ItemBuilder() interface{} {
	return &Item{}
}

func main() {
	flag.Parse()
	args := flag.Args()
	log.Println(args)
	if len(args) < 1 {
		log.Println("starting url not specified")
		os.Exit(1)
	}

	qName := "item-queue"
	qDir := "tmp"
	segmentSize := 50
	var queue *dque.DQue
	var err error

	if _, err := os.Stat(qDir); os.IsNotExist(err) {
		createError := os.Mkdir(qDir, os.ModePerm)
		if createError != nil {
			panic(createError)
		}
		// Create a new queue with segment size of 50
		queue, err = dque.New(qName, qDir, segmentSize, ItemBuilder)
		if err != nil {
			log.Println(err)
		}
		// Add an item to the queue
		err = queue.Enqueue(&Item{"args[0]", 1})
		if err != nil {
			log.Println(err)
		}
	} else {
		// You can reconsitute the queue from disk at any time
		queue, err = dque.Open(qName, qDir, segmentSize, ItemBuilder)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("queue loading success")
	//go func() { queue.Enqueue(&Item{args[0], 1}) }()

	var iface interface{}
	if iface, err = queue.Peek(); err != nil {
		if err != dque.ErrEmpty {
			log.Fatal("Error peeking at item ", err)
		}
	}

	// Dequeue the next item in the queue
	if iface, err = queue.Dequeue(); err != nil {
		if err != dque.ErrEmpty {
			log.Fatal("Error dequeuing item ", err)
		}
	}

	item, ok := iface.(*Item)
	if !ok {
		log.Fatal("Dequeued object is not an Item pointer")
		log.Println(ok)
	}
	log.Println(item)

	//for uri := range queue {
	//	log.Println("log item:",uri)
	//	entity, err := enqueue(uri, queue)
	//	if err != nil {
	//		log.Println("enqueue error:", err.Error(), uri)
	//	}
	//	log.Println(entity.ImageURL)
	//	_, err = request_handlers.CreateEntity(entity)
	//	log.Println("entity added", entity.Title)
	//	if err != nil {
	//		log.Println(err.Error(), uri)
	//	}
	//}
}

//func enqueue(uri string, queue dque.DQue) (models.Entity, error) {
//	log.Println("fetching", uri)
//	visited[uri] = true
//
//	var (
//		entity models.Entity
//		err    error
//		body   *html.Node
//	)
//
//	entity = models.Entity{}.SetSource(uri).SetSourceSignature("trusted")
//
//	resp, err := request_handlers.GetRequest(uri)
//	if err != nil {
//		return entity, err
//	}
//
//	doc, err := libraries.HTMLStringToDoc(resp)
//	if err != nil {
//		return entity, err
//	}
//
//	entity.Title, body, err = parsers.ParseHTMLContent(doc)
//	if err != nil {
//		return entity, err
//	}
//
//	//clean html code by removing unwanted information
//	htmlCleaner := clean_html.HtmlCleaner{Config: clean_html.Config{
//		LineBreakers:   []string{"div", "caption"},
//		IgnoreElements: []string{"noscript", "script", "style", "input"},
//		IgnoreStrings:  []string{"[", "]", "edit", "Jump to search", "Jump to navigation"},
//		IgnoreTitles:   []string{"(page does not exist)", ":"},
//		IgnoreClasses:  []string{"box-Multiple_issues"},
//	}}
//	result, linkedEntities, imageList, defaultImageSource := htmlCleaner.CleanHTML(uri, body)
//	entity.ImageURL = defaultImageSource
//
//	// queue new links for crawling
//	for _, linkedEntity := range linkedEntities {
//		if !visited[linkedEntity.GetSource()] {
//			go func(url string) {
//				queue <- url
//			}(linkedEntity.GetSource())
//		}
//		entity = entity.AddLink(models.Link{}.SetTitle(linkedEntity.GetTitle()).AddDate(entity.GetSourceDate()))
//	}
//
//	for _, image := range imageList {
//		go func(payload models.Upload) {
//			request_handlers.UploadImage(payload)
//		}(image)
//	}
//
//	// save linkedEntities (create empty if not exist)
//	entity, err = request_handlers.AddEntitiesAsLinks(entity, linkedEntities)
//	entity = entity.SetAttribute("content", models.Value{
//		ValueType:   "html",
//		ValueString: result,
//	}).AddCategory("Wikipedia")
//	return entity, nil
//}
