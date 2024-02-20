package handlers

import (
	"testing"

	repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
	"github.com/golang-jwt/jwt/v5"
)

func TestLoginAPI(t *testing.T) {

  // Setup test data
  email := "test@example.com"
  pwd := "password123"

  // Initialise handler
  userRepo := &repo.UserRepository{}
  handler := NewLoginHandler(userRepo)

  // Valid credentials
  user, token, err := handler.Login(email, pwd)
  
  // Assert
  if err != nil {
    t.Fatal(err)
  }

  if user.Email != email {
    t.Errorf("email mismatch") 
  }

  // Verify token
  err = ValidateToken(token)
  if err != nil {
      t.Fatal(err)
  } 

  // Optionally check claims match 
  claims := jwt.MapClaims{}
  jwt.ParseWithClaims(token, claims, VerifyKeyFunc)

  if claims["id"] != user.ID {
      t.Error("claim mismatch")
  }
}

func TestLoginAPIInvalidCreds(t *testing.T) {

  // Wrong password
  email := "test@example.com"
  pwd := "wrongpass"

  // Execute handler
  userRepo := &repo.UserRepository{}
  handler := NewLoginHandler(userRepo)  
  user, token, err := handler.Login(email, pwd)

  // Assert
  if user != nil || token != "" {
    t.Error("expected empty response") 
  }

  if err != ErrInvalidCredentials {
    t.Error("expected invalid credentials error")
  }

}
