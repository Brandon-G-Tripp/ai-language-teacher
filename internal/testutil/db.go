package testutil

import (
	"log"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitTestDB() *gorm.DB {
    env.LoadEnv()

    var db *gorm.DB
    var err error
    db, err = database.ConnectDB("test")
    if err != nil {
        panic("Failed to connect to database: %v" + err.Error())
    } 

    // Enable logger for test
    db.Logger.LogMode(logger.Info)

    // Run Migrations 
    err = database.Migrate("test")
    if err != nil {
        log.Fatalf("Error in test database migration: %v", err)
    } 

    return db
}

func CloseTestDB(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        panic("Failed to get SQL DB connection: " + err.Error())
    } 

    sqlDB.Close()
} 
