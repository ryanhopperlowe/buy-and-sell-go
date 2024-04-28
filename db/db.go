package db

import (
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func MigrateDB() {
	DB.AutoMigrate(&model.Listing{})
	DB.AutoMigrate(&model.User{})
}
