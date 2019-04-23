package controllers

import (
	"backend-rest/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveIndex(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/workers/scraping_index", SaveScrapingIndex)

	body := models.ScrapingIndex{DateLastNew: time.Now(), DateScraping: time.Now(), ScraperID: "test2", DeviceID: "deviceTest"}
	json, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", "/api/workers/scraping_index", bytes.NewBuffer(json))
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func TestGetIndex(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/api/workers/scraping_index", GetScrapingIndex)

	request, _ := http.NewRequest("GET", "/api/workers/scraping_index?scraper_id=test", nil)
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}
