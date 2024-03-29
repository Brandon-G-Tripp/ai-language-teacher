import (
    "bytes"
    "encoding/json"
    "net/http"
    "testing"

    "github.com/Brandon-G-Tripp/ai-language-teacher/internal/testutil"

    handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
)

func TestCreateConversation(t *testing.T) {
    // Test cases
    testCases := []struct {
        name string
        input handler_models.CreateConversationRequest
        expectedStatus int 
        expectedError string
    } {
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
            // Arrange 
            ctx := testutil.GetTestGinContext()
            data, _ := json.Marshal(tc.input)
            req, _ := http.NewRequest(http.MethodPost, "/api/v1/conversations", bytes.NewReader(data))
            ctx.Request = req


            // Act 
            conversation, err := conversationHandler.CreateConversation(tc.input)

            // Assert 
            if tc.expectedError != "" {
                if err == nil || err.Error() != tc.expectedError {
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
            if got, want := ctx.Writer.Status(), tc.expectedStatus; got != want {
                t.Errorf("Unexpected status code: got %v want %v", got, want)
            } 
        } 
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
            // Arrange
            ctx := testutil.GetTestGinContext()
            req, _ := http.NewRequest(http.MethodGet, "/api/v1/conversations/1", nil)
            ctx.Request = req

            // Act
            conversation, err := conversationHandler.GetConversation(tc.conversationID)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.Error() != tc.expectedError {
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
            if got, want := ctx.Writer.Status(), tc.expectedStatus; got != want {
                t.Errorf("Unexpected status code: got %v want %v", got, want)
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
            // Arrange
            ctx := testutil.GetTestGinContext()
            data, _ := json.Marshal(tc.input)
            req, _ := http.NewRequest(http.MethodPut, "/api/v1/conversations/1", bytes.NewReader(data))
            ctx.Request = req

            // Act
            conversation, err := conversationHandler.UpdateConversation(tc.conversationID, tc.input)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.Error() != tc.expectedError {
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
            if got, want := ctx.Writer.Status(), tc.expectedStatus; got != want {
                t.Errorf("Unexpected status code: got %v want %v", got, want)
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
            // Arrange
            ctx := testutil.GetTestGinContext()
            req, _ := http.NewRequest(http.MethodDelete, "/api/v1/conversations/1", nil)
            ctx.Request = req

            // Act
            err := conversationHandler.DeleteConversation(tc.conversationID)

            // Assert
            if tc.expectedError != "" {
                if err == nil || err.Error() != tc.expectedError {
                    t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
            }
            if got, want := ctx.Writer.Status(), tc.expectedStatus; got != want {
                t.Errorf("Unexpected status code: got %v want %v", got, want)
            }
        })
    }
}
