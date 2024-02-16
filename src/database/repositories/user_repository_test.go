package repositories

import (
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
) 

func TestCreateUser(t *testing.T) {
    // Arrange 
    db, err := database.ConnectDB("test")
    if err != nil {
        t.Fatalf("Error connecting to the database: %v", err)
    } 

    // deferred cleanup
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

    // init User Repo
    repo := NewUserRepository(db)

    user := models.User{
        Name: "John Doe",
        Email: "john@doe.com",
        Password: "password",
    } 

    // Act

    err = repo.Create(&user)

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
