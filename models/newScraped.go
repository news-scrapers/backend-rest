package models

import (
	u "backend-rest/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type NewScraped struct {
	Headline  string    `json:"headline"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	NewsPaper string    `json:"newspaper"`
	ScraperID string    `json:"scraper_id" bson:"scraper_id"`
}

func (newScraped *NewScraped) Create() map[string]interface{} {

	db := GetDB()
	collection := db.Collection("NewsScraped")

	options := options.FindOneAndReplaceOptions{}
	upsert := true
	options.Upsert = &upsert
	err := collection.FindOneAndReplace(context.Background(), bson.M{"url": newScraped.Url}, newScraped, &options)

	//_, err := collection.InsertOne(context.Background(), newScraped)

	if err == nil {
		resp := u.Message(true, "success")
		resp["new"] = newScraped
		return resp
	} else {
		//fmt.Println(err)
		return nil
	}

}
