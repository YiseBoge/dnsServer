package models

import (
	"dnsServer/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type DomainName struct {
	gorm.Model

	Name     string
	Address  string
	LastRead time.Time
}

func CreateDomainName(model DomainName) {
	database := db.GetOpenDatabase()
	var prev DomainName
	database.Where("name = ? AND address = ?", model.Name, model.Address).First(&prev)
	if prev.ID <= 0 {
		database.Save(&model)
	}
	defer database.Close()
}

func UpdateDomainName(model DomainName) {
	database := db.GetOpenDatabase()
	database.Model(&model).Update(&model)
	defer database.Close()
}

func AllDomainNames() []DomainName {
	var models []DomainName
	database := db.GetOpenDatabase()
	database.Find(&models)
	defer database.Close()
	return models
}

func DomainNameById(id int) DomainName {
	var model DomainName
	database := db.GetOpenDatabase()
	database.First(&model, id)
	defer database.Close()
	return model
}

func DomainNamesByName(name string) []DomainName {
	var models []DomainName
	database := db.GetOpenDatabase()
	database.Find(&models, "name = ?", name)
	defer database.Close()
	return models
}

func DomainNamesByAddress(address string) []DomainName {
	var models []DomainName
	database := db.GetOpenDatabase()
	database.Find(&models, "address = ?", address)
	defer database.Close()
	return models
}

func DeleteDomainName(model DomainName) DomainName {
	database := db.GetOpenDatabase()
	database.Delete(&model)
	defer database.Close()
	return model
}

func MigrateDomainName() {
	database := db.GetOpenDatabase()
	database.AutoMigrate(&DomainName{})
	defer database.Close()
}
