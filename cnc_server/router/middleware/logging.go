package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// LoggingMiddleware function log for program
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var schema string
		if request.TLS != nil {
			schema = "https://"
		}
		schema = "http://"

		URL := schema + request.Host + request.RequestURI
		record := fmt.Sprintf("%s %s %s", request.RemoteAddr, URL, request.UserAgent())
		log.Printf("%v", record)
		next.ServeHTTP(response, request)
	})
}
