package database

import (
	"github.com/Kaa-dan/skill-management/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	var err error
	//db call
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	//db err
	if err != nil {
		panic("failed to connect databse")
	}

	DB.AutoMigrate(&models.UserModel{})
}
