package models

import "time"

type ServerNode struct {
	Address    string
	Port       string
	Descriptor string
}

type ServerModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Address       string
	Port          string
	ParentAddress string
	ParentPort    string
}
