package main

import (
	"backend-rest/controllers"
	"backend-rest/middlewares"
	"io"
	log "log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const logPath = "logs.log"

var Logger *log.Logger

func main() {

	InitLogger()

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
	log.Println("Starting server on port " + port)
	log.Println(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func InitLogger() {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file")
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

}
