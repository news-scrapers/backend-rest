package controllers

import (
	"backend-rest/models"
	u "backend-rest/utils"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var AddNew = func(w http.ResponseWriter, r *http.Request) {

	NewScraped := &models.NewScraped{}

	err := json.NewDecoder(r.Body).Decode(NewScraped)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := NewScraped.Create()
	u.Respond(w, resp)
}

var AddMany = func(w http.ResponseWriter, r *http.Request) {

	newScraped := []models.NewScraped{}

	err := json.NewDecoder(r.Body).Decode(&newScraped)
	ms := fmt.Sprintf("inserting %v news from ", len(newScraped)) + newScraped[0].NewsPaper
	log.Info(ms)

	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	_ = models.CreateManyNewsScraped(newScraped)
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}
