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
	//defer database.Close()
}

func (DomainName) FindAll(database *gorm.DB) []DomainName {
	var models []DomainName
	database.Find(&models)
	//defer database.Close()
	return models
}

func (DomainName) FindById(database *gorm.DB, id int) DomainName {
	var model DomainName
	database.First(&model, id)
	//defer database.Close()
	return model
}

func (DomainName) FindByName(database *gorm.DB, name string) []DomainName {
	var models []DomainName
	database.Find(&models, "name = ?", name)
	//defer database.Close()
	return models
}

func (DomainName) FindByAddress(database *gorm.DB, address string) []DomainName {
	var models []DomainName
	database.Find(&models, "address = ?", address)
	//defer database.Close()
	return models
}

func (model *DomainName) Delete(database *gorm.DB) {
	database.Delete(model)
	//defer database.Close()
}

func (DomainName) Migrate() {
	database := db.GetOpenDatabase()
	cacheDatabase := db.GetOpenCacheDatabase()

	database.AutoMigrate(&DomainName{})
	cacheDatabase.AutoMigrate(&DomainName{})

	defer database.Close()
	defer cacheDatabase.Close()
}
