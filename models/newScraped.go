package models

import (
	u "backend-rest/utils"
	"context"
	"time"

	log "log"

	"go.mongodb.org/mongo-driver/bson"
)

type NewScraped struct {
	Page      int       `json:"page"`
	FullPage  bool      `json:"full_page"`
	Headline  string    `json:"headline"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	NewsPaper string    `json:"newspaper"`
	ScraperID string    `json:"scraper_id" bson:"scraper_id"`
	ID        string    `json:"id"`
}

func (newScraped *NewScraped) Create() map[string]interface{} {

	db := GetDB()
	collection := db.Collection("NewsContentScraped")

	// options := options.FindOneAndReplaceOptions{}
	// upsert := true
	// options.Upsert = &upsert
	// err := collection.FindOneAndReplace(context.Background(), bson.M{"url": newScraped.Url}, newScraped, &options)
	result := &NewScraped{}
	err := collection.FindOne(context.Background(), bson.M{"url": newScraped.Url}).Decode(result)

	if err != nil {
		err2, _ := collection.InsertOne(context.Background(), newScraped)
		if err2 != nil {
			log.Println("saved new with " + newScraped.Url)
			resp := u.Message(true, "success")
			resp["new"] = newScraped
			return resp
		} else {
			log.Println("error saving new with url " + newScraped.Url)
			log.Println(err)
			return nil
		}
	} else {
		log.Println("record already exists with url " + newScraped.Url)
		return nil
	}

	//_, err := collection.InsertOne(context.Background(), newScraped)

}

func CreateManyNewsScraped(newsScraped []NewScraped) error {

	db := GetDB()
	collection := db.Collection("NewsContentScraped")
	docs := []interface{}{}

	for _, result := range newsScraped {
		docs = append(docs, result)
	}
	_, err := collection.InsertMany(context.Background(), docs)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
