package handlers

import (
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
)

func TestUserAuthFlowIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    } 
    // Test data 
    email := "testauthflow@example.com"
    pwd := "password123"

    // 1. Call Signup handler 
    user, _, err := signUpHandler.SignUp(models.SignUpRequest{
        Name: "Johnny Auth Test",
        Email: email,
        Password: pwd,
    })

    if err != nil {
        t.Fatalf("SignUp failed: %v", err)
    } 

    // 2. Call login handler with user creds
    loggedInUser, token_login, err := loginHandler.Login(email, pwd)
    if err != nil {
        t.Fatalf("Login failed: %v", err)
    } 

    // Assert that the login was successful
    if token_login == "" {
        t.Error("Expected a token to be returned after successful login")
    } 

    // Assert that the logged in user matches the created user
    if loggedInUser.ID != user.ID || loggedInUser.Email != user.Email {
        t.Errorf("Logged in user data does not match created user")
    } 

    // 3. Call logut handler with token 
    err = logoutHandler.Logout(token_login)
    if err != nil {
        t.Errorf("Logout failed: %v", err)
    } 

    // 4. Assert that the token is now invalid
    err = authService.ValidateToken(token_login)
    if err == nil {
        t.Error("Expected token to be invalid after logout, but it is still valid")
    } else if err != auth.ErrInvalidToken {
        t.Errorf("Expected ErrInvalidToken, but got: %v", err)
    } 
} 
