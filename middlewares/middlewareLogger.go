package middlewares

import (
	"fmt"
	"net/http"
)

var MiddlewareLogger = func (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
		return
	})
}

