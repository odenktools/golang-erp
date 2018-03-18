package models

import (
	"github.com/jinzhu/gorm"
)

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=moeloet dbname=edan_golang sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	db.Debug().AutoMigrate(&Company{})
	return db, nil
}
