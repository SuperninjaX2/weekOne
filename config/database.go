package config

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"weekOne/models"
)

var Db *gorm.DB

func Database() {
	var err error
	Db, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		return
	}
}
