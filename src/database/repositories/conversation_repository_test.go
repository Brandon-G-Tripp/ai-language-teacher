package repositories

import (
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
) 

func TestCreateConversation(t *testing.T) {
    tx := db.Begin()

    repo := NewConversationRepository(db)
    conversation := &models.Conversation{
        UserID: 1,
        Title: "Test Conversation",
    }

    err := repo.Create(conversation)

    if err != nil {
        t.Errorf("Error creating conversation: %v", err)
    } 
    if conversation.ID == 0 {
        t.Error("Conversation ID not set after creation")
    } 

    tx.Rollback()
} 

func TestGetConversationById(t *testing.T) {
    // Arrange
    repo := NewConversationRepository(db)
    conversation := &models.Conversation{
        UserID: 1,
        Title:  "Test Conversation",
    }
    repo.Create(conversation)

    // Act
    fetchedConversation, err := repo.GetById(conversation.ID)

    // Assert
    if err != nil {
        t.Errorf("Error getting conversation by ID: %v", err)
    }
    if fetchedConversation.ID != conversation.ID || fetchedConversation.Title != conversation.Title {
        t.Error("Fetched conversation does not match")
    }
}

func TestUpdateConversation(t *testing.T) {
    // Arrange
    repo := NewConversationRepository(db)
    conversation := &models.Conversation{
        UserID: 1,
        Title:  "Test Conversation",
    }
    repo.Create(conversation)

    // Act
    conversation.Title = "Updated Conversation"
    err := repo.Update(conversation)

    // Assert
    if err != nil {
        t.Errorf("Error updating conversation: %v", err)
    }
    fetchedConversation, _ := repo.GetById(conversation.ID)
    if fetchedConversation.Title != "Updated Conversation" {
        t.Error("Conversation title not updated")
    }
}

func TestDeleteConversation(t *testing.T) {
    // Arrange
    repo := NewConversationRepository(db)
    conversation := &models.Conversation{
        UserID: 1,
        Title:  "Test Conversation",
    }
    repo.Create(conversation)

    // Act
    err := repo.Delete(conversation)

    // Assert
    if err != nil {
        t.Errorf("Error deleting conversation: %v", err)
    }
    _, err = repo.GetById(conversation.ID)
    if err == nil {
        t.Error("Conversation still exists after deletion")
    }
}
