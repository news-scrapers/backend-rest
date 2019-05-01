package models

import (
	u "backend-rest/utils"
	"context"
	"fmt"
	"time"
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

	_, err := collection.InsertOne(context.Background(), newScraped)

	if err == nil {
		resp := u.Message(true, "success")
		resp["new"] = newScraped
		return resp
	} else {
		fmt.Println(err)
		return nil
	}

}
