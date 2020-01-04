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

func (model *DomainName) Save() {
	database := db.GetOpenDatabase()

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
	defer database.Close()
}

func (DomainName) FindAll() []DomainName {
	var models []DomainName
	database := db.GetOpenDatabase()
	database.Find(&models)
	defer database.Close()
	return models
}

func (DomainName) FindById(id int) DomainName {
	var model DomainName
	database := db.GetOpenDatabase()
	database.First(&model, id)
	defer database.Close()
	return model
}

func (DomainName) FindByName(name string) []DomainName {
	var models []DomainName
	database := db.GetOpenDatabase()
	database.Find(&models, "name = ?", name)
	defer database.Close()
	return models
}

func (DomainName) FindByAddress(address string) []DomainName {
	var models []DomainName
	database := db.GetOpenDatabase()
	database.Find(&models, "address = ?", address)
	defer database.Close()
	return models
}

func (model *DomainName) Delete() {
	database := db.GetOpenDatabase()
	database.Delete(model)
	defer database.Close()
}

func (DomainName) Migrate() {
	database := db.GetOpenDatabase()
	database.AutoMigrate(&DomainName{})
	defer database.Close()
}
