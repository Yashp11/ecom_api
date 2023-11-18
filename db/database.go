package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"myapi/models"
)

var db *gorm.DB

// Init initializes the database connection
func Init() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// AutoMigrate the models
	db.AutoMigrate(&models.Mobile{})

	return db, nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
