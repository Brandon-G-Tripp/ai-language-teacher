package database

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestDatabaseConnection(t *testing.T) {
    db, err := connectDB("../../.env.test")
    if err != nil {
        t.Errorf("Failed to connect to database: %v", err)
    } 

    // Enable logger for test
    db.Logger.LogMode(logger.Info)

    sqlDB, err := db.DB()
    defer func() {
        if err != nil {
            t.Errorf("Failed to close database connection: %v", err)
        } 

        err = sqlDB.Close()
        if err != nil {
            t.Errorf("Failed to close the database connection: %v", err)
        } 
    }()

    err = sqlDB.Ping()
    if err != nil {
        t.Errorf("Failed to ping the database: %v", err)
    } 
} 
