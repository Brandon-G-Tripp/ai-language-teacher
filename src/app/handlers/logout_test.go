package handlers

import (
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

func TestLogoutHandlerWithValidToken(t *testing.T) {
    // Setup 
    authService := auth.NewAuthService()

    userRepo :=  user_repo.NewUserRepository(database.DB)
    hashedPwd, err := authService.HashPassword("logouttest")
    if err != nil {
        t.Fatal("Error returned from hashing password")
    }
    user := db_models.User{
        Name: "John Doe Logout",
        Email: "testlogout@mail.com",
        Password: hashedPwd,
    } 

    err = userRepo.Create(&user)
    if err != nil {
        t.Fatalf("Error creating user: %v", err)
    } 

    token, err := authService.GenerateJWT(&user)
    if err != nil {
        t.Fatalf("Error generating token: %v", err)
    } 

    // Execute 
    handler := NewLogoutHandler(userRepo, authService)
    err = handler.Logout(token)

    // Assert 
    if err != nil {
        t.Errorf("Logout failed: %v", err)
    } 

    err = authService.ValidateToken(token)
    if err == nil {
        t.Error("Expected token to be invalidated, but it is still valid")
    }
} 


func TestLogoutHandlerWithInvalidToken(t *testing.T) {
    // Setup
    userRepo := user_repo.NewUserRepository(database.DB)
    authService := auth.NewAuthService()

    // Execute
    handler := NewLogoutHandler(userRepo, authService)
    err := handler.Logout("invalid_token")

    // Assert 
    if err == nil {
        t.Error("Expected error for invalid token, but got nil")
    } else {
        expectedErr := auth.ErrInvalidToken
        if err != expectedErr {
            t.Errorf("Expected error '%v', but got '%v'", expectedErr, err)
        } 
    } 
} 
