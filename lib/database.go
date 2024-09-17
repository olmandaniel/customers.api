package lib

import (
	"github.com/olman99/customers-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123456 dbname=customers_db port=5432 sslmode=disable TimeZone=America/Lima"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Customer{})

	return db, nil
}
