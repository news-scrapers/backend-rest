package main

import (
	"backend-rest/controllers"
	"backend-rest/middlewares"
	"io"
	ioutil "io/ioutil"
	log "log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const logPath = "logs.log"

var Logger *log.Logger

func main() {

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

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
	Info.Println("Starting server on port " + port)
	Error.Println(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file")
	}

	Logger = log.New(file,
		"PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
