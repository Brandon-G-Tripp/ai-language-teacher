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

    db.Exec("TRUNCATE users")
    os.Exit(exitCode)

} 

func TestCreateUser(t *testing.T) {
    // Arrange 
    // start transaction
    tx := db.Begin()

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

    // Rollbak user created
    tx.Rollback()

} 

func TestGetUserById(t *testing.T) {
    db.Exec("TRUNCATE users")
    tx := db.Begin()

    repo := NewUserRepository(db)

    user := models.User{
        Name: "John Doe",
        Email: "john@doe.com",
        Password: "password",
    }

    err := repo.Create(&user)
    if err != nil {
        t.Fatalf("Error creating user: %v", err)
    } 

    // Get the user by ID
    fetchedUser, err := repo.GetById(user.ID)
    if err != nil {
        t.Fatalf("Error getting user by ID: %v", err)
    } 

    // Verif the fetchyed user matches
    if fetchedUser.ID != user.ID || fetchedUser.Name != user.Name || fetchedUser.Email != user.Email {
        t.Errorf("Fetched user does not match")
    }

    tx.Rollback()
} 


func TestUpdateUser(t *testing.T) {
    // Start transaction
    db.Exec("TRUNCATE users")
    tx := db.Begin()

    repo := NewUserRepository(db)

    // Create user 
    user := models.User{
        Name: "John Doe",
        Email: "john@doe.com",
        Password: "password",
    }    
    err := repo.Create(&user)
    if err != nil {
        t.Fatalf("Error creating user: %v", err)
    }

    // Update the user
    user.Name = "Jane Doe"
    err = repo.Update(&user)
    if err != nil {
        t.Fatalf("Error updating user: %v", err)
    }

    // Fetch the user
    fetchedUser, err := repo.GetById(user.ID)
    if err != nil {
        t.Fatalf("Error getting user by ID: %v", err)
    }

    // Verify the name was updated
    if fetchedUser.Name != "Jane Doe" {
        t.Errorf("Expected name to be updated")
    }

    tx.Rollback()
} 

func TestDeleteUser(t *testing.T) {
    db.Exec("TRUNCATE users")
    tx := db.Begin()

    repo := NewUserRepository(db)
    // Create a user
    user := models.User{
        Name: "John Doe",
        Email: "john@doe.com",
        Password: "password",
    }    
    err := repo.Create(&user)
    if err != nil {
        t.Fatalf("Error creating user: %v", err)
    }

    // Delete the user
    err = repo.Delete(&user)
    if err != nil {
        t.Fatalf("Error deleting user: %v", err)
    }
    tx.Commit()

    // Try to fetch the user
    _, err = repo.GetById(user.ID)

    // Verify the user was deleted
    if err == nil {
        t.Errorf("Expected error fetching deleted user")
    }

    tx.Rollback()
}

func TestGetAllUsers(t *testing.T) {
    db.Exec("TRUNCATE users")
    tx := db.Begin()

    repo := NewUserRepository(db)
    // Create some users
    user1 := models.User{
        Name: "John Doe",
        Email: "john@doe.com",
        Password: "password",
    }
    user2 := models.User{
        Name: "Jane Doe",
        Email: "jane@doe.com",
        Password: "password",
    }

    err := repo.Create(&user1)
    if err != nil {
        t.Fatalf("Error creating user 1: %v", err)
    }

    err = repo.Create(&user2)
    if err != nil {
        t.Fatalf("Error creating user 2: %v", err)
    }

    // Get all users
    users, err := repo.GetAll()
    if err != nil {
        t.Fatalf("Error getting all users: %v", err)
    }

    // Verify we got the 2 users back
    if len(users) != 2 {
        t.Errorf("Expected 2 users, got %d", len(users))
    }

    // Verify the users match
    if (users[0].ID != user1.ID || users[0].Name != user1.Name || users[0].Email != user1.Email) && 
       (users[1].ID != user2.ID || users[1].Name != user2.Name || users[1].Email != user2.Email) {
        t.Errorf("Fetched users do not match")
    }

    tx.Rollback()
}

