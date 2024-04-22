package handlers

import (
    "testing"
    "net/http"

    handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
)

func TestCreateMessage(t *testing.T) {
    testCases := []struct {
        name string
        input handler_models.CreateMessageRequest
        expectedStatus int
        expectedError string
    } {
        {
            name: "Valid request", 
            input: handler_models.CreateMessageRequest{
                ConversationID: 1,
                UserID: 1,
                Content: "Test Message",
            },
            expectedStatus: http.StatusCreated,
            expectedError: "",
        },
        {
            name: "Invalid request - Missing ConversationID", 
            input: handler_models.CreateMessageRequest{
                UserID: 1,
                Content: "Test Message",
            },
            expectedStatus: http.StatusBadRequest,
            expectedError: "Conversation ID is required",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            message, err := messageHandler.CreateMessage(tc.input)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != tc.expectedError {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if message == nil {
                    t.Error("Expected message, but got nil")
                }
                // Add more assertions for message fields if necessary
            }
        })
    }
}

func TestGetMessagesByConversationID(t *testing.T) {
    // Test cases
    testCases := []struct {
        name           string
        conversationID uint
        expectedStatus int
        expectedError  string
    }{
        {
            name:           "Valid request",
            conversationID: 1,
            expectedStatus: http.StatusOK,
            expectedError:  "",
        },
        {
            name:           "Invalid request - Conversation not found",
            conversationID: 999,
            expectedStatus: http.StatusNotFound,
            expectedError:  "Conversation not found",
        },
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            messages, err := messageHandler.GetMessagesByConversationID(tc.conversationID)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != "Conversation not found" {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if messages == nil {
                    t.Error("Expected messages, but got nil")
                }
                // Add more assertions for messages if needed
            }
        })
    }
}

func TestUpdateMessage(t *testing.T) {
    // Test cases
    testCases := []struct {
        name          string
        messageID     uint
        input         handler_models.UpdateMessageRequest
        expectedStatus int
        expectedError string
    }{
        {
            name:      "Valid request",
            messageID: 1,
            input: handler_models.UpdateMessageRequest{
                Content: "Updated Message",
            },
            expectedStatus: http.StatusOK,
            expectedError:  "",
        },
        {
            name:      "Invalid request - Message not found",
            messageID: 999,
            input: handler_models.UpdateMessageRequest{
                Content: "Updated Message",
            },
            expectedStatus: http.StatusNotFound,
            expectedError:  "Message not found",
        },
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            message, err := messageHandler.UpdateMessage(tc.messageID, tc.input)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != tc.expectedError {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if message == nil {
                    t.Error("Expected message, but got nil")
                }
                // Add more assertions for message fields if needed
            }
        })
    }
}

func TestDeleteMessage(t *testing.T) {
    // Test cases
    testCases := []struct {
        name           string
        messageID      uint
        expectedStatus int
        expectedError  string
    }{
        {
            name:           "Valid request",
            messageID:      1,
            expectedStatus: http.StatusNoContent,
            expectedError:  "",
        },
        {
            name:           "Invalid request - Message not found",
            messageID:      999,
            expectedStatus: http.StatusNotFound,
            expectedError:  "Message not found",
        },
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            _, err := messageHandler.DeleteMessage(tc.messageID)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != tc.expectedError {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
            }
        })
    }
} 
