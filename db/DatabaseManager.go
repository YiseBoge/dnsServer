package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

func GetOpenDatabase() *gorm.DB {
	workingDirectory, _ := os.Getwd()
	db, err := gorm.Open("sqlite3", workingDirectory+"/db/dns_database.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetOpenCacheDatabase() *gorm.DB {
	workingDirectory, _ := os.Getwd()
	db, err := gorm.Open("sqlite3", workingDirectory+"/db/cache_dns_database.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
