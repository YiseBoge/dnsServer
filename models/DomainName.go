package models

import (
	"dnsServer/db"
	"fmt"
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

func (model *DomainName) Save(database *gorm.DB) {
	var prev DomainName
	database.Where("name = ? AND address = ?", model.Name, model.Address).First(&prev)
	model.LastRead = time.Now()
	if model.ID == 0 && prev.ID == 0 {
		fmt.Println(model.ID)
		database.Create(model)
	} else if model.ID > 0 {
		database.First(&prev, model.ID)
		database.Model(&prev).Update(model)
	}
}

func (DomainName) SaveAll(database *gorm.DB, domainNames []DomainName) {
	for _, domainName := range domainNames {
		d := DomainName{Name: domainName.Address, Address: domainName.Address}
		d.Save(database)
	}
}

func (DomainName) FindAll(database *gorm.DB) []DomainName {
	var models []DomainName
	database.Find(&models)
	return models
}

func (DomainName) FindById(database *gorm.DB, id int) DomainName {
	var model DomainName
	database.First(&model, id)
	return model
}

func (DomainName) FindByName(database *gorm.DB, name string) []DomainName {
	var models []DomainName
	database.Find(&models, "name = ?", name)
	return models
}

func (DomainName) FindByAddress(database *gorm.DB, address string) []DomainName {
	var models []DomainName
	database.Find(&models, "address = ?", address)
	return models
}

func (model *DomainName) Delete(database *gorm.DB) {
	database.Delete(model)
}

func (DomainName) Migrate() {
	database := db.GetOpenDatabase()
	cacheDatabase := db.GetOpenCacheDatabase()

	database.AutoMigrate(&DomainName{})
	cacheDatabase.AutoMigrate(&DomainName{})

	defer database.Close()
	defer cacheDatabase.Close()
}

func ClearTimedOut(timeout int) {
	database := db.GetOpenCacheDatabase()
	all := DomainName{}.FindAll(database)
	for _, domain := range all {
		difference := int(time.Now().Sub(domain.LastRead).Hours())
		if difference > timeout {
			domain.Delete(database)
		}
	}
	database.Close()
}
