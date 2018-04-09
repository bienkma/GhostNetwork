package handler

import (
	"log"
	"fmt"
	"net/http"
)

// Function log middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var schema string = "http://"
		if request.TLS != nil {
			schema = "https://"
		}

		URL := schema + request.Host + request.RequestURI
		record := fmt.Sprintf("%s %s %s", request.RemoteAddr, URL, request.UserAgent())
		log.Printf("%v", record)
		next.ServeHTTP(response, request)
	})
}
