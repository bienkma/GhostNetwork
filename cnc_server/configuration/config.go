package configuration

import (
	"github.com/caarlos0/env"
)

// CNCCfg struct init variable
type CNCCfg struct {
	CNCAddress   string `env:"CNCAddress" envDefault:"0.0.0.0"`
	DBAddress    string `env:"DBAddress" envDefault:"127.0.0.1"`
	DBName       string `env:"DBName" envDefault:"cnc"`
	DBPort       int    `env:"DBPort" envDefault:"3306"`
	DBUser       string `env:"DBUser" envDefault:"root"`
	DBPass       string `env:"DBPass" envDefault:"changeme"`
	DBLogMode    bool   `env:"DBLogMode" envDefault:"false"`
	DBMaxConn    int    `env:"DBMaxConn" envDefault:"1024"`
	CNCPathStore string `env:"CNCPathStore" envDefault:"/cnc"`
	USEHTTPs     bool   `env:"USEHttps" envDefault:"true"`
	SSLCert      string `env:"SSLCert" envDefault:"/etc/cnc/server.crt"`
	SSLKey       string `env:"SSLKey" envDefault:"/etc/cnc/server.key"`
}

// CNCCfg function get variable
func CNCConfig() (CNCCfg, error) {
	var cfg = CNCCfg{}
	err := env.Parse(&cfg)
	return cfg, err
}
