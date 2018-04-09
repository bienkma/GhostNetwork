package database

type Bots struct {
	GUID        string `gorm:"size:32; primary_key; unique"`
	CLIENTIP    string `gorm:"not null"`
	WHOAMI      string `gorm:"not null"`
	OS          string `gorm:"not null"`
	INSTALL     string `gorm:"not null"`
	ADMIN       string `gorm:"not null"`
	AV          string
	CPU         string `gorm:"not null"`
	GPU         string
	VERSION     string `gorm:"not null"`
	SYSINFO     string `gorm:"not null"`
	LASTCHECKIN string `gorm:"not null"`
}

type Accounts struct {
	UUID     string `gorm:"size:32; primary_key; unique"`
	NAME     string `gorm:"not null"`
	PASSWORD string `gorm:"not null"`
	EMAIL    string `gorm:"not null"`
	ISSUPPER bool   `gorm:"not null"`
}
