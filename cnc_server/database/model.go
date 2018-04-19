package database

// Bots model
type Bots struct {
	GUID        string `gorm:"size:32; primary_key; unique"`
	WHOAMI      string `gorm:"not null"`
	ADMIN       string `gorm:"not null"`
	OS          string `gorm:"not null"`
	VERSION     string `gorm:"not null"`
	INSTALLED   string `gorm:"not null"`
	AV          string
	IPLAN       string `gorm:"not null"`
	IPWAN       string `gorm:"not null"`
	CPU         string `gorm:"not null"`
	GPU         string
	SYSINFO     string `gorm:"not null"`
	LASTCHECKIN string `gorm:"not null"`
}

// Accounts model
type Accounts struct {
	UUID     string `gorm:"size:32; primary_key; unique"`
	NAME     string `gorm:"not null"`
	PASSWORD string `gorm:"not null"`
	EMAIL    string `gorm:"not null"`
	ISADMIN  bool   `gorm:"not null"`
}
