package controllers

import (
	"backend-rest/models"
	u "backend-rest/utils"
	"encoding/json"
	"net/http"
)

var SaveScrapingIndex = func(w http.ResponseWriter, r *http.Request) {

	NewIndex := &models.ScrapingIndex{}

	err := json.NewDecoder(r.Body).Decode(NewIndex)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := NewIndex.Save()
	u.Respond(w, resp)
}

var GetScrapingIndex = func(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	scraperID := v.Get("scraper_id")
	index := models.GetCurrentIndex(scraperID)
	var resp map[string]interface{}
	resp["data"] = index
	u.Respond(w, resp)
}
