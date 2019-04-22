package controllers

import (
	"backend-rest/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNew(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/workers/new_scraped", AddNew)

	body := models.NewScraped{Headline: "New"}
	json, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", "/api/workers/new_scraped", bytes.NewBuffer(json))
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}
