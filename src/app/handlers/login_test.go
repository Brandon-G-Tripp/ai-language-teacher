package handlers

import (
	"log"
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
	"github.com/golang-jwt/jwt/v5"
)

func TestLoginAPI(t *testing.T) {
  // Setup test data
  email := "testapi@example.com"
  pwd := "password123"

  // Initialise handler
    hashPassword, err := authService.HashPassword(pwd)
    if err != nil {
        t.Fatal("Error returned from hashing password")
    }
    userToCreate := db_models.User{
        Name: "John Doe",
        Email: email,
        Password: hashPassword,
    } 

    // log.Printf("User to create: %v", userToCreate)
  err = userRepo.Create(&userToCreate)

    // Assert
  if err != nil {
     t.Fatalf("Error creating the user: %v", err)
  } 

  handler := NewLoginHandler(userRepo, authService)

  // Valid credentials
  user, token, err := handler.Login(email, pwd)
  log.Printf("User in test valid: %v", user)

  // Assert
  if err != nil {
      t.Fatalf("Err is present in login_test: %v", err)
  }

  if user.Email != email {
    t.Errorf("email mismatch") 
  }


  // Verify token
  err = authService.ValidateToken(token)
  if err != nil {
      t.Fatalf("Error verifying token: %v", err)
  } 

  parsed, err := jwt.Parse(token, authService.VerifyKeyFunc) 

  claims, ok := parsed.Claims.(jwt.MapClaims)
  if !ok {
      t.Fatal("invalid claims")
  } 

  id := claims["id"].(float64)

  if id != float64(user.ID) {
      t.Errorf("invalid id claim")
  }
}

func TestLoginAPIInvalidCreds(t *testing.T) {

  // Wrong password
  email := "test@example.com"
  pwd := "wrongpass"

  // Execute handler
  authService := auth.NewAuthService()

  // Initialise handler
    hashPassword, err := authService.HashPassword("password123")
    if err != nil {
        t.Fatal("Error returned from hashing password")
    }
  userRepo := user_repo.NewUserRepository(database.DB)
    userToCreate := db_models.User{
        Name: "John Doe2",
        Email: email,
        Password: hashPassword,
    } 

  err = userRepo.Create(&userToCreate)

    // Assert
  if err != nil {
     t.Fatalf("Error creating the user: %v", err)
  } 
  handler := NewLoginHandler(userRepo, authService)  
  _, _, err = handler.Login(email, pwd)

  // Assert
  if err == nil {
    t.Error("expected error response") 
  }

  if err != ErrInvalidCredentials {
    t.Error("expected invalid credentials error but did not receive")
  }

}
