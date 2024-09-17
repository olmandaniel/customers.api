package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint   `json:"id"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Country   string `json:"country"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (Customer) TableName() string {
	return "customers"
}
