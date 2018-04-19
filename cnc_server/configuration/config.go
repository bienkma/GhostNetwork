package configuration

import (
	"fmt"
	"os"
)

// CNCConfig struct init variable
type CNCConfig struct {
	CNCAddress   string
	DBAddress    string
	DBName       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBLogMode    string
	DBMaxConn    string
	CNCPathStore string
	USEHttps     string
	SSLCert      string
	SSLKey       string
}

// CNCConf function get variable
func CNCConf() (conf CNCConfig, err error) {
	conf.CNCAddress = os.Getenv("CNC_ADDRESS")
	conf.DBAddress = os.Getenv("DB_ADDRESS")
	conf.DBName = os.Getenv("DB_NAME")
	conf.DBPort = os.Getenv("DB_PORT")
	conf.DBUser = os.Getenv("DB_USER")
	conf.DBPass = os.Getenv("DB_PASS")
	conf.DBLogMode = os.Getenv("DB_LOG_MODE")
	conf.DBMaxConn = os.Getenv("DB_MAX_CONN")
	conf.CNCPathStore = os.Getenv("CNC_PATH_STORE")

	if len(conf.CNCAddress) == 0 || len(conf.CNCPathStore) == 0 || len(conf.DBAddress) == 0 {
		err := fmt.Errorf("Not set environment")
		return conf, err
	}
	return conf, nil
}
