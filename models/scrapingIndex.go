package models

import (
	u "backend-rest/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type ScrapingIndex struct {
	DateLastNew     time.Time `json:"date_last_new" bson:"date_last_new"`
	DateScraping    time.Time `json:"date_scraping" bson:"date_scraping"`
	LastHistoricUrl string    `json:"last_historic_url" bson:"last_historic_url"`
	NewsPaper       string    `json:"newspaper" bson:"newspaper"`
	ScraperID       string    `json:"scraper_id" bson:"scraper_id"`
	DeviceID        string    `json:"device_id" bson:"device_id"`
}

func (scrapingIndex *ScrapingIndex) Save() map[string]interface{} {

	db := GetDB()
	collection := db.Collection("ScrapingIndex")
	options := options.FindOneAndReplaceOptions{}
	upsert := true
	options.Upsert = &upsert

	err := collection.FindOneAndReplace(context.Background(), bson.M{"scraper_id": scrapingIndex.ScraperID}, scrapingIndex, &options)

	if err == nil {
		resp := u.Message(true, "success")
		resp["new"] = scrapingIndex
		return resp
	} else {
		//_, err2 := collection.InsertOne(context.Background(), scrapingIndex)
		fmt.Println("error updating")
		fmt.Println(err)
		return nil
	}

}

func GetCurrentIndex(scraperID string) (scrapingIndex ScrapingIndex) {
	db := GetDB()
	collection := db.Collection("ScrapingIndex")

	options := options.FindOptions{}
	// Sort by `_id` field descending
	options.Sort = bson.D{{"date_last_new", 1}}
	limit := int64(1)
	options.Limit = &limit

	results := []ScrapingIndex{}
	res, err := collection.Find(context.Background(), bson.M{"scraper_id": scraperID}, &options)
	res.Decode(&results)
	if err != nil {
		panic(err)
	}
	scrapingIndex = results[0]
	return scrapingIndex

}
