package database

import (
	"github.com/jinzhu/gorm"

	"github.com/bienkma/GhostNetwork/cnc_server/configuration"
)

var seeder *BotDB

type BotDB struct {
	Db     *gorm.DB
	Config *configuration.CNCCfg
}

// Build container db for pool connection
func Builder(db *gorm.DB, config *configuration.CNCCfg) {
	seeder = &BotDB{Db: db, Config: config}
}
