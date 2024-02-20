package handlers

import (
    "errors"

    auth "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
    repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
    db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
) 

var (
  ErrInvalidCredentials = errors.New("invalid email or password")
  ErrUserNotFound = errors.New("user not found")
)

type LoginHandler struct {
  UserRepo repo.UserRepository
}

func NewLoginHandler(repo repo.UserRepository) *LoginHandler {
  return &LoginHandler{UserRepo: repo}
}

func (h *LoginHandler) Login(email, password string) (*db_models.User, string, error) {

  // Find user
  user, err := h.UserRepo.GetByEmail(email)
  if err != nil {
    return nil, "", err
  }

  // Not found
  if user == nil {
    return nil, "", ErrUserNotFound
  }

  authService := auth.NewAuthService()

  // Check password
  if authService.ValidatePassword(user.Password, password) {
    return nil, "", ErrInvalidCredentials
  }

  // Generate JWT token
  token, err := authService.GenerateJWT(user)
  if err != nil {
    return nil, "", err
  }

  return user, token, nil
}
