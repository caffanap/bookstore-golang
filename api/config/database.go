package config

import (
	"caffanap/bookstore/api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase(url string) {
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}
	Database = db
}

func MigrationDatabase() {
	// register models
	book := &models.Book{}

	Database.AutoMigrate(
		book,
	)
}
