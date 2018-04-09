package configuration

import (
	"os"
	"encoding/json"
)

type CNCConfig struct {
	CNCAddress      string
	CNCPort         int
	DBAddress       string
	DBName          string
	DBPort          int
	DBUsername      string
	DBPassword      string
	DBLogMode       bool
	DBMaxConnection int
	Storage string
}

func CNCConf() CNCConfig {

	file, err := os.Open("/etc/cnc/cnc.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	cfg := CNCConfig{}

	if err := decoder.Decode(&cfg); err != nil {
		panic(err)
	}
	return cfg
}