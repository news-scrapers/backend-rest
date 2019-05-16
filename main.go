package main

import (
	"backend-rest/controllers"
	"backend-rest/middlewares"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const logPath = "logs.log"

func main() {

	configLogs()

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/workers/new_scraped", controllers.AddNew).Methods("POST")
	router.HandleFunc("/api/workers/new_scraped_many", controllers.AddMany).Methods("POST")
	router.HandleFunc("/api/workers/scraping_index", controllers.SaveScrapingIndex).Methods("POST")
	router.HandleFunc("/api/workers/scraping_index", controllers.GetScrapingIndex).Methods("GET")
	router.HandleFunc("/api/workers/scraping_index_newspaper", controllers.GetScrapingIndexNewsPaper).Methods("GET")

	// router.Use(middlewares.JwtAuthentication) //attach JWT auth middleware
	router.Use(middlewares.MiddlewareLogger) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Info("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func configLogs() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	//log.SetOutput(f)
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}
