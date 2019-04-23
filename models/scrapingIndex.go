package models

import (
	u "backend-rest/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		resp["data"] = scrapingIndex
		return resp
	} else {
		_, errDelete := collection.DeleteOne(context.Background(), scrapingIndex)
		_, errInsert := collection.InsertOne(context.Background(), scrapingIndex)
		//fmt.Println("error updating")
		//fmt.Println(err)
		fmt.Println(errInsert)
		fmt.Println(errDelete)
		return nil
	}

}

func GetCurrentIndex(scraperID string) (scrapingIndex ScrapingIndex) {
	db := GetDB()
	collection := db.Collection("ScrapingIndex")

	options := options.FindOneOptions{}
	// Sort by `_id` field descending
	options.Sort = bson.D{{"date_last_new", int32(1)}}

	results := ScrapingIndex{}
	err := collection.FindOne(context.Background(), bson.M{"scraper_id": scraperID}, &options).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	scrapingIndex = results
	return scrapingIndex

}
