package handlers

import (
    "net/http"
    "testing"

    handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
)

func TestCreateConversation(t *testing.T) {
    // Test cases
    testCases := []struct {
        name string
        input handler_models.CreateConversationRequest
        expectedStatus int 
        expectedError string
    }{
        {
            name: "Valid request",
            input: handler_models.CreateConversationRequest{
                UserID: 1,
                Title:  "Test Conversation",
            },
            expectedStatus: http.StatusCreated,
            expectedError:  "",
        },
        {
            name: "Invalid request - Missing UserID",
            input: handler_models.CreateConversationRequest{
                Title: "Test Conversation",
            },
            expectedStatus: http.StatusBadRequest,
            expectedError:  "User ID is required",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act 
            conversation, err := conversationHandler.CreateConversation(tc.input)

            // Assert 
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != tc.expectedError {
                    t.Errorf("Expected error '%s', but got %v'", tc.expectedError, err)
                } 
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                } 
                if conversation == nil {
                    t.Error("Expected conversation, but got nil")
                } 
                // add more assertions for convo fields if necessary
            } 
        }) 
    } 
} 


func TestGetConversation(t *testing.T) {
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
            conversation, err := conversationHandler.GetConversation(tc.conversationID)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != tc.expectedError {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if conversation == nil {
                    t.Error("Expected conversation, but got nil")
                }
                // Add more assertions for conversation fields if needed
            }
        })
    }
}

func TestUpdateConversation(t *testing.T) {
    // Test cases
    testCases := []struct {
        name           string
        conversationID uint
        input          handler_models.UpdateConversationRequest
        expectedStatus int
        expectedError  string
    }{
        {
            name:           "Valid request",
            conversationID: 1,
            input: handler_models.UpdateConversationRequest{
                Title: "Updated Conversation",
            },
            expectedStatus: http.StatusOK,
            expectedError:  "",
        },
        {
            name:           "Invalid request - Conversation not found",
            conversationID: 999,
            input: handler_models.UpdateConversationRequest{
                Title: "Updated Conversation",
            },
            expectedStatus: http.StatusNotFound,
            expectedError:  "Conversation not found",
        },
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            conversation, err := conversationHandler.UpdateConversation(tc.conversationID, tc.input)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.(handler_models.ApiError).Message != tc.expectedError {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if conversation == nil {
                    t.Error("Expected conversation, but got nil")
                }
                // Add more assertions for conversation fields if needed
            }
        })
    }
}

func TestDeleteConversation(t *testing.T) {
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
            expectedStatus: http.StatusNoContent,
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
            _, err := conversationHandler.DeleteConversation(tc.conversationID)

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
