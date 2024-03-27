package repositories

import (
    "testing"

    "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

func TestCreateMessage(t *testing.T) {
    // Arrange
    repo := NewMessageRepository(db)
    message := &models.Message{
        ConversationID: 1,
        UserID:         1,
        Content:        "Test Message",
    }

    // Act
    err := repo.Create(message)

    // Assert
    if err != nil {
        t.Errorf("Error creating message: %v", err)
    }
    if message.ID == 0 {
        t.Error("Message ID not set after creation")
    }
}

func TestGetMessageById(t *testing.T) {
    // Arrange
    repo := NewMessageRepository(db)
    message := &models.Message{
        ConversationID: 1,
        UserID:         1,
        Content:        "Test Message",
    }
    repo.Create(message)

    // Act
    fetchedMessage, err := repo.GetById(message.ID)

    // Assert
    if err != nil {
        t.Errorf("Error getting message by ID: %v", err)
    }
    if fetchedMessage.ID != message.ID || fetchedMessage.Content != message.Content {
        t.Error("Fetched message does not match")
    }
}

func TestUpdateMessage(t *testing.T) {
    // Arrange
    repo := NewMessageRepository(db)
    message := &models.Message{
        ConversationID: 1,
        UserID:         1,
        Content:        "Test Message",
    }
    repo.Create(message)

    // Act
    message.Content = "Updated Message"
    err := repo.Update(message)

    // Assert
    if err != nil {
        t.Errorf("Error updating message: %v", err)
    }
    fetchedMessage, _ := repo.GetById(message.ID)
    if fetchedMessage.Content != "Updated Message" {
        t.Error("Message content not updated")
    }
}

func TestDeleteMessage(t *testing.T) {
    // Arrange
    repo := NewMessageRepository(db)
    message := &models.Message{
        ConversationID: 1,
        UserID:         1,
        Content:        "Test Message",
    }
    repo.Create(message)

    // Act
    err := repo.Delete(message)

    // Assert
    if err != nil {
        t.Errorf("Error deleting message: %v", err)
    }
    _, err = repo.GetById(message.ID)
    if err == nil {
        t.Error("Message still exists after deletion")
    }
}
