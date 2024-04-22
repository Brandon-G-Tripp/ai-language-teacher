package database_test

import (
	"os"
	"testing"

	// "github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/internal/testutil"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
)

func TestMain(m *testing.M) {
    // Init DB
    test_db := testutil.InitTestDB()
    database.DB = test_db
    // Setup handler
    // run tests
    exitCode := m.Run()

    // Close connection 
    testutil.CloseTestDB(database.DB)

    os.Exit(exitCode)
} 

// func init() {
//     env.LoadEnv()

//     var err error
//     db, err = ConnectDB("test")
//     if err != nil {
//         panic("Failed to connect to database: %v" + err.Error())
//     } 

//     // Enable logger for test
//     db.Logger.LogMode(logger.Info)

//     // Run Migrations 
//     err = Migrate("test")
//     if err != nil {
//         log.Fatalf("Error in test database migration: %v", err)
//     } 
// } 

// func TestMain(m *testing.M) {
//     // run tests
//     exitCode := m.Run()

//     // Close connection 
//     sqlDB, err := db.DB()
//     if err != nil {
//         panic("Failed to get SQL DB connection: " + err.Error())
//     } 
//     defer sqlDB.Close()

//     os.Exit(exitCode)

// } 

func TestDatabaseConnection(t *testing.T) {
    // Check if we can ping db
    sqlDB, err := database.DB.DB()
    if err != nil {
        t.Errorf("Failed to get SQL DB connection in TestDatabase test: %v", err)
    } 

    err = sqlDB.Ping()
    if err != nil {
        t.Errorf("Failed to ping the database: %v", err)
    } 
} 
