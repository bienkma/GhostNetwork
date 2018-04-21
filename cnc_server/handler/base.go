package handler

import (
	"github.com/bienkma/GhostNetwork/cnc_server/configuration"
	"github.com/jinzhu/gorm"
)

// BotServer include into handler
type BotServer struct {
	Db  *gorm.DB
	Cfg configuration.CNCCfg
}

// UserAgent for Bot Server check
const UserAgent string = "FLDSH3LFDBVSLfdvfdLDYeuglghjgndghioerGH"
