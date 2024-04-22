package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"github.com/gin-gonic/gin"
)

var db *gorm.DB 

func TestMain(m *testing.M) {
    env.LoadEnv()

    db, err := database.ConnectDB("test")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    } 

    // Enable logger for test
    db.Logger.LogMode(logger.Info)

    // Run Migrations 
    err = database.Migrate("test")
    if err != nil {
        log.Fatalf("Error in test database migration: %v", err)
    } 

    sqlDB, err := db.DB()

    // Run tests
    m.Run()

    defer sqlDB.Close()

    os.Exit(0)
} 

func TestAPI(t *testing.T) {
	// Set up Gin router
	r := gin.Default()
	setupRoutes(r, db)

	// Test signup endpoint
	t.Run("Signup", func(t *testing.T) {
		signupData := handler_models.SignUpRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		jsonData, _ := json.Marshal(signupData)

		req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		// Additional assertions for the response body can be added here
	})

	// Test login endpoint
	t.Run("Login", func(t *testing.T) {
		loginData := handler_models.LoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		jsonData, _ := json.Marshal(loginData)

		req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		// Additional assertions for the response body can be added here
	})

	// Test protected endpoints (e.g., conversation and message endpoints)
	t.Run("ProtectedEndpoints", func(t *testing.T) {
		// Perform login to get an access token
		loginData := handler_models.LoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		jsonData, _ := json.Marshal(loginData)

		loginReq, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonData))
		loginReq.Header.Set("Content-Type", "application/json")

		loginW := httptest.NewRecorder()
		r.ServeHTTP(loginW, loginReq)

		var loginResp handler_models.LoginResponse
		json.Unmarshal(loginW.Body.Bytes(), &loginResp)

		// Use the access token for subsequent requests
		accessToken := loginResp.Token

		// Test conversation endpoints
		conversationData := handler_models.CreateConversationRequest{
			Title: "Test Conversation",
		}
		jsonData, _ = json.Marshal(conversationData)

		conversationReq, _ := http.NewRequest("POST", "/api/v1/conversations", bytes.NewBuffer(jsonData))
		conversationReq.Header.Set("Content-Type", "application/json")
		conversationReq.Header.Set("Authorization", "Bearer "+accessToken)

		conversationW := httptest.NewRecorder()
		r.ServeHTTP(conversationW, conversationReq)

		if conversationW.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, conversationW.Code)
		}
		// Additional assertions for the response body can be added here

		// Test message endpoints
		messageData := handler_models.CreateMessageRequest{
			Content: "Test Message",
		}
		jsonData, _ = json.Marshal(messageData)

		messageReq, _ := http.NewRequest("POST", "/api/v1/conversations/1/messages", bytes.NewBuffer(jsonData))
		messageReq.Header.Set("Content-Type", "application/json")
		messageReq.Header.Set("Authorization", "Bearer "+accessToken)

		messageW := httptest.NewRecorder()
		r.ServeHTTP(messageW, messageReq)

		if messageW.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, messageW.Code)
		}
		// Additional assertions for the response body can be added here
	})
}

