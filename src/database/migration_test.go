package database

import (
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

func TestMigrations(t *testing.T) {
    // use db connection established in TestMain in orm_test.go
    if DB == nil {
        t.Fatal("Database connection is nil")
    } 

    if !DB.Migrator().HasTable(&models.User{}) {
        t.Error("User table does not exist")
    } 

    if !DB.Migrator().HasTable(&models.Conversation{}) {
        t.Error("Conversation table does not exist")
    } 

    if !DB.Migrator().HasTable(&models.Message{}) {
        t.Error("Message table does not exist")
    } 

    // Check if foreign key constraints are properly defined
    conversation := models.Conversation{Title: "Test Conversation"}
    DB.Create(&conversation)

    message := models.Message{
        ConversationID: conversation.ID,
        UserID: 1,
        Content: "Test Message",
    }
    err := DB.Create(&message).Error
    if err != nil {
        t.Errorf("Error creating message: %v", err)
    } 

    // Clean up the test data
    DB.Delete(&message)
    DB.Delete(&conversation)
} 
