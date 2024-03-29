package handlers

import (
	"errors"

	auth "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
) 

var (
  ErrInvalidCredentials = errors.New("invalid email or password")
  ErrUserNotFound = errors.New("user not found")
)

type LoginHandler struct {
  UserRepo *user_repo.UserRepository
  AuthService *auth.AuthService
}

func NewLoginHandler(repo *user_repo.UserRepository, authService *auth.AuthService) *LoginHandler {
    return &LoginHandler{UserRepo: repo, AuthService: authService}
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


  // Check password
  if !h.AuthService.ValidatePassword(user.Password, password) {
    return nil, "", ErrInvalidCredentials
  }

  // Generate JWT token
  token, err := h.AuthService.GenerateJWT(user)
  if err != nil {
    return nil, "", err
  }

  return user, token, nil
}
