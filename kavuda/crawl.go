package main

import (
	"Kavuda/ada_derana"
	"Kavuda/ceylon_today"
	"Kavuda/daily_mirror"
	"Kavuda/daily_news"
	"Kavuda/models"
	"Kavuda/the_island"
	"Kavuda/utils"
	"log"
)

func main() {

	crawl(ada_derana.AdaDeranaDecoder{})
	crawl(ceylon_today.CeylonTodayDecoder{})
	crawl(daily_mirror.DailyMirrorDecoder{})
	crawl(daily_news.DailyNewsDecoder{})
	crawl(the_island.TheIslandDecoder{})
}

func crawl(decoder models.IDecoder) {
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
		newsItem, contentString, err := decoder.FillNewsContent(newsItem)
		if err != nil {
			panic(err)
		}

		log.Println("		Reading News Article Completed.")

		//decode to entity
		entity := utils.EntityFromNews(newsItem, decoder.GetSourceTitle()).SetSourceSignature("trusted")

		//save entity with NER processing
		utils.ProcessAndSaveEntity(entity, contentString)
	}
}
