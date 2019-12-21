package models

import (
	u "backend-rest/utils"
	"context"
	"time"

	log "log"

	"go.mongodb.org/mongo-driver/bson"
)

type NewScraped struct {
	Page      int       `json:"page" bson:"page"`
	FullPage  bool      `json:"full_page" bson:"full_page"`
	Headline  string    `json:"headline" bson:"headline"`
	Date      time.Time `json:"date" bson:"date"`
	Content   string    `json:"content" bson:"content"`
	Url       string    `json:"url" bson:"url"`
	NewsPaper string    `json:"newspaper" bson:"newspaper"`
	ScraperID string    `json:"scraper_id" bson:"scraper_id"`
	ID        string    `json:"id" bson:"id"`
	Tags      []string  `json:"tags" bson:"tags"`
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
