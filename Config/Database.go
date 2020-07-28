package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	User     string
	Password string
	Host     string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		User:     "mydb58",
		Password: "Tz24zr-c~lHw",
		Host:     "den1.mssql7.gear.host",
		DBName:   "mydb58",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s&connection+timeout=30",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DBName,
	)
}
