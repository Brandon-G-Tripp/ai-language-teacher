package main

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm/logger"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
)


func TestMain(m *testing.M) {
    env.LoadEnv()

    db, err := database.ConnectDB("test")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    } 

    // Enable logger for test
    db.Logger.LogMode(logger.Info)

    // Run Migrations 
    err = database.Migrate("test")
    if err != nil {
        log.Fatalf("Error in test database migration: %v", err)
    } 

    sqlDB, err := db.DB()

    // Run tests
    m.Run()

    defer sqlDB.Close()

    os.Exit(0)
} 
