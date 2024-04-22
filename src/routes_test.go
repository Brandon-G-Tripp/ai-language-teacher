package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
)

func TestRoutes(t *testing.T) {
    // Set up Gin router
    r := gin.Default()
    setupRoutes(r, db)

    // Test public routes
    t.Run("Public Routes", func(t *testing.T) {
        // Test signup route
        signupData := models.SignUpRequest{
            Email:    "test@example.com",
            Password: "password",
        }
        signupBody, _ := json.Marshal(signupData)
        signupReq, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(signupBody))
        signupReq.Header.Set("Content-Type", "application/json")
        signupResp := httptest.NewRecorder()
        r.ServeHTTP(signupResp, signupReq)
        if signupResp.Code != http.StatusOK {
            t.Errorf("Expected status code %d, got %d", http.StatusOK, signupResp.Code)
        }

        // Test login route
        loginData := models.LoginRequest{
            Email:    "test@example.com",
            Password: "password",
        }
        loginBody, _ := json.Marshal(loginData)
        loginReq, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(loginBody))
        loginReq.Header.Set("Content-Type", "application/json")
        loginResp := httptest.NewRecorder()
        r.ServeHTTP(loginResp, loginReq)
        if loginResp.Code != http.StatusOK {
            t.Errorf("Expected status code %d, got %d", http.StatusOK, loginResp.Code)
        }
    })

    // Test protected routes
    t.Run("Protected Routes", func(t *testing.T) {
        // Perform login to get an access token
        loginData := models.LoginRequest{
            Email:    "test@example.com",
            Password: "password",
        }
        loginBody, _ := json.Marshal(loginData)
        loginReq, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(loginBody))
        loginReq.Header.Set("Content-Type", "application/json")
        loginResp := httptest.NewRecorder()
        r.ServeHTTP(loginResp, loginReq)

        var loginResponse models.LoginResponse
        json.Unmarshal(loginResp.Body.Bytes(), &loginResponse)
        accessToken := loginResponse.Token

        // Test conversation routes
        conversationData := models.CreateConversationRequest{
            Title: "Test Conversation",
        }
        conversationBody, _ := json.Marshal(conversationData)
        conversationReq, _ := http.NewRequest("POST", "/api/v1/conversations", bytes.NewBuffer(conversationBody))
        conversationReq.Header.Set("Content-Type", "application/json")
        conversationReq.Header.Set("Authorization", "Bearer "+accessToken)
        conversationResp := httptest.NewRecorder()
        r.ServeHTTP(conversationResp, conversationReq)
        if conversationResp.Code != http.StatusOK {
            t.Errorf("Expected status code %d, got %d", http.StatusOK, conversationResp.Code)
        }

        // Test message routes
        messageData := models.CreateMessageRequest{
            Content: "Test Message",
        }
        messageBody, _ := json.Marshal(messageData)
        messageReq, _ := http.NewRequest("POST", "/api/v1/conversations/1/messages", bytes.NewBuffer(messageBody))
        messageReq.Header.Set("Content-Type", "application/json")
        messageReq.Header.Set("Authorization", "Bearer "+accessToken)
        messageResp := httptest.NewRecorder()
        r.ServeHTTP(messageResp, messageReq)
        if messageResp.Code != http.StatusOK {
            t.Errorf("Expected status code %d, got %d", http.StatusOK, messageResp.Code)
        }
    })
}
