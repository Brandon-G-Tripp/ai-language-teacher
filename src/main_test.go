package main

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm/logger"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"github.com/joho/godotenv"
)


func TestMain(m *testing.M) {
    err := godotenv.Load(".env.test")
    if err != nil {
        log.Fatalf("Failed to load environment variables: %v", err)
    } 

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
    defer sqlDB.Close()

    // Run tests
    exitVal := m.Run()


    // Close Connection
    sqlDB.Close()

    os.Exit(exitVal)
} 
