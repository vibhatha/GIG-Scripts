package main

import (
	"GIG-Scripts/kavuda/ada_derana"
	"GIG-Scripts/kavuda/ceylon_today"
	"GIG-Scripts/kavuda/daily_mirror"
	"GIG-Scripts/kavuda/daily_news"
	"GIG-Scripts/kavuda/models"
	"GIG-Scripts/kavuda/the_island"
	"GIG-Scripts/kavuda/utils"
	"log"
	"sync"
)

func main() {
	decoders := []models.IDecoder{
		ada_derana.AdaDeranaDecoder{},
		ceylon_today.CeylonTodayDecoder{},
		daily_mirror.DailyMirrorDecoder{},
		daily_news.DailyNewsDecoder{},
		the_island.TheIslandDecoder{},
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(decoders))

	for _, decoder := range decoders {
		go crawl(decoder, &waitGroup)
	}

	waitGroup.Wait()

}

func crawl(decoder models.IDecoder, wg *sync.WaitGroup) {
	//extract news items from site
	newsItems, err := decoder.ExtractNewsItems()
	if err != nil {
		panic(err)
	}
	log.Println("News links extracted...")
	log.Println(len(newsItems), "news items found\n ")

	//for each news article
	log.Println("Reading News...")
	for _, newsItem := range newsItems {

		log.Println("	Item: ", newsItem.Title)
		log.Println("	News: ", newsItem.Link)
		log.Println("	Date: ", newsItem.Date)
		newsItem, contentString, err := decoder.FillNewsContent(newsItem)
		if err != nil {
			panic(err)
		}

		log.Println("		Reading News Article Completed.")

		//decode to entity
		entity := utils.EntityFromNews(newsItem, decoder.GetSourceTitle()).SetSourceSignature("trusted").
			AddCategories(newsItem.Categories).SetSource(newsItem.Link)

		//save entity with NER processing
		utils.ProcessAndSaveEntity(entity, contentString)
	}
	wg.Done()
}
