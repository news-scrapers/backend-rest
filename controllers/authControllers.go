package controllers

import (
	"backend-rest/models"
	u "backend-rest/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(*account)
	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp, code := models.Login(account.Email, account.Password)
	if code == 401 {
		w.WriteHeader(http.StatusUnauthorized)
	} else if code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	u.Respond(w, resp)
}
