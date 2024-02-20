package auth

import (
	"testing"

	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

func TestGenerateToken(t *testing.T) {
    user := db_models.User{ID: 1, Email: "test@email.com"}

    token, err := generateToken(&user)
    if err != nil {
        t.Errorf("GenerateToken returned error: %v", err)
    } 

    // Validate token is not empyt
    if token == "" {
        t.Error("Empty token generated")
    } 
} 

func TestValidateToken(t *testing.T) {
    // Create user
    user := db_models.User{ID: 1, Email: "test@email.com"}

    // Generate token 
    token, err := generateToken(&user)
    if err != nil {
        t.Errorf("Failed to generate token: %v", err)
        return 
    }

    // Validate token
    err = validateToken(token)
    if err != nil {
        t.Errorf("ValidateToken returned error: %v", err)
    }

    // Test with an invalid token 
    err = validateToken("")
    if err == nil {
        t.Error("Expected an error for an invalid token, but got nil")
    }
} 
