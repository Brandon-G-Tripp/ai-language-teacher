package repositories

import (
	"os"
	"testing"

	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
) 

var db *gorm.DB

func init() {
    env.LoadEnv()

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
} 

func TestMain(m *testing.M) {
    // run tests
    exitCode := m.Run()

    // Close connection 
    sqlDB, err := db.DB()
    if err != nil {
        panic("Failed to get SQL DB connection: " + err.Error())
    } 
    defer sqlDB.Close()

    os.Exit(exitCode)

} 

func TestCreateUser(t *testing.T) {
    // Arrange 

    // deferred cleanup
    // init User Repo
    repo := NewUserRepository(db)

    user := models.User{
        Name: "John Doe",
        Email: "john@doe.com",
        Password: "password",
    } 

    // Act

    err := repo.Create(&user)

    // Assert
    if err != nil {
        t.Fatalf("Error creating the user: %v", err)
    } 

    // verify user was created
    var createdUser *models.User
    createdUser, err = repo.GetByEmail(user.Email)
    if err != nil {
        t.Fatalf("Error getting user by email: %v", err)
    } 

    if createdUser.ID == 0 {
        t.Errorf("Expected non-zero ID for created user")
    } 

} 
