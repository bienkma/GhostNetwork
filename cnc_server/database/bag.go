package database

import (
	"github.com/jinzhu/gorm"

	"github.com/bienkma/GhostNetwork/cnc_server/configuration"
)

var seeder *BotDB

type BotDB struct {
	Db     *gorm.DB
	Config *configuration.CNCConfig
}

// Build container db for pool connection
func Builder(db *gorm.DB, config *configuration.CNCConfig) {
	seeder = &BotDB{Db: db, Config: config}
}
