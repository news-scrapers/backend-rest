package models

import (
	u "backend-rest/utils"
	"context"
	"time"
)

type NewScraped struct {
	Headline  string    `json:"headline"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	NewsPaper string    `json:"newspaper"`
	ScraperID string    `json:"scraper_id"`
}

func (newScraped *NewScraped) Create() map[string]interface{} {

	db := GetDB()
	collection := db.Collection("NewsScraped")
	_, err := collection.InsertOne(context.Background(), newScraped)

	if err == nil {
		resp := u.Message(true, "success")
		resp["new"] = newScraped
		return resp
	} else {
		return nil
	}

}
