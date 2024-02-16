package database

import (
	"testing"

)

func TestDatabaseConnection(t *testing.T) {
    db, err := ConnectDB("test")
    if err != nil {
        t.Errorf("Failed to connect to database: %v", err)
    } 

    // Check if we can ping db
    sqlDB, err := db.DB()
    err = sqlDB.Ping()
    if err != nil {
        t.Errorf("Failed to ping the database: %v", err)
    } 
} 
