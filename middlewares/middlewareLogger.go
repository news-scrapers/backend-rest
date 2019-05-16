package middlewares

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var MiddlewareLogger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		date := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
		log.Info(date + " -->" + " " + r.Method + " " + r.URL.Path + " " + r.URL.RawQuery)
		next.ServeHTTP(w, r)
		return
	})
}
