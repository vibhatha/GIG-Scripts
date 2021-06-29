package main

import (
	"GIG-SDK"
	"GIG-SDK/enums/ValueType"
	"GIG-SDK/models"
	"GIG-SDK/request_handlers"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

/**
config before running
 */
var categories = []string{"Press Releases"}

/**
input a web url containing pdf files
automatically download and create entities for all of pdf files
 */
var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
// if your img's are properly formed with doublequotes then use this, it's more efficient.
// var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)
func findImages(htm string) []string {
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

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
	images := findImages(resp)

	for _, img := range images {
		var entities []models.Entity
		if strings.Contains(img, "Release") && strings.Contains(img, "/images/") {
			imageUrl := "https://www.dgi.gov.lk" + img
			releaseDate, err := time.Parse("2006.01.02", strings.Split(img,"/")[2])
			if err != nil {
				log.Println(err)
				panic("invalid filename")
			}
			textContent, err := request_handlers.GetRequest(config.OCRServer + imageUrl)
			if err != nil {
				panic(err)
			}
			//NER extraction
			entityTitles, err := request_handlers.ExtractEntityNames(textContent)
			if err != nil {
				log.Println(err)
			}
			title := strings.Replace(img, "/images/", "", -1)
			title = strings.Replace(title, "/", "_", -1)

			entity := models.Entity{
				Title: title,
			}.
				SetSource("Press Release " + img).
				SetSourceDate(releaseDate).
				SetSourceSignature("trusted").
				AddCategories(categories).
				SetAttribute("extracted_text", models.Value{
					ValueType:   "string",
					ValueString: textContent,
					Date:        releaseDate,
					Source:      "Press Release " + img,
					UpdatedAt:   time.Now(),
				})

			fmt.Println(entity)
			fmt.Println(entityTitles)

			for _, mentionedEntity := range entityTitles {
				childEntity := models.Entity{}.
					SetTitle(models.Value{ValueType: ValueType.String, ValueString: mentionedEntity.GetEntityName(), Source: "Press Release " + img, Date: releaseDate}).
					SetSource("Press Release " + img).
					SetSourceDate(releaseDate).
					AddCategory(mentionedEntity.GetCategory()).AddLink(models.Link{}.SetTitle(title).AddDate(releaseDate)).
					SetAttribute("source", models.Value{
						ValueType:   "string",
						ValueString: "Press Release",
						Date:        releaseDate,
						Source:      "Press Release " + img,
						UpdatedAt:   time.Now(),
					})

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
