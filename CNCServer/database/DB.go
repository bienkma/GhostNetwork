package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type MySqlConfig struct {
	Host          string
	Port          int
	DbName        string
	DBUsername    string
	Password      string
	LogEnable     bool
	MaxConnection int
}

func (dbConf MySqlConfig) GetMySqlUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
		dbConf.DBUsername, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName)
}

func OpenMySQL(conf MySqlConfig) *gorm.DB {
	db, ok := gorm.Open("mysql", conf.GetMySqlUrl())
	if ok != nil {
		log.Fatal(ok)
	}
	db.DB().SetMaxIdleConns(0)
	// Init Table
	bot := Bots{}
	account := Accounts{}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	//db.DropTableIfExists(bot, account)
	db.AutoMigrate(bot, account)
	// ForeignKey
	// db.Model(&account).AddForeignKey("uuid", "bot(GUID)", "CASCADE", "RESTRICT")

	db.LogMode(conf.LogEnable)
	db.DB().SetMaxOpenConns(conf.MaxConnection)
	return db
}
