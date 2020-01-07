package db

import (
	"dnsServer/config"
	"dnsServer/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"time"
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

func ClearTimedOut() {
	db := GetOpenCacheDatabase()
	all := models.DomainName{}.FindAll(db)
	for _, domain := range all {
		difference := int(time.Now().Sub(domain.LastRead).Hours())
		if difference > config.LoadConfig().Timeout {
			domain.Delete(db)
		}
	}
	db.Close()
}
