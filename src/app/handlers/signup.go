package handlers

import (
	"net/http"

	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

type SignUpHandler struct {
    UserRepo *user_repo.UserRepository
    AuthService *auth.AuthService
} 

func NewSignUpHandler(repo *user_repo.UserRepository, authService *auth.AuthService) *SignUpHandler {
    return &SignUpHandler{
        UserRepo: repo,
        AuthService: authService,
    }
} 

func (h *SignUpHandler) SignUp(req handler_models.SignUpRequest) (*db_models.User, string, error) {

    // Validate Request
    if !IsEmailValid(req.Email) {
        return nil, "", handler_models.ApiError{
            Message: "Invalid email address",
            Code: 400,
        }
    } 
    
    user, err := h.UserRepo.GetByEmail(req.Email)

    // Email already exists check
    if err == nil && user.ID != 0 {
        return nil, "", handler_models.ApiError{
            Message: "Internal server error",
            Code: http.StatusInternalServerError,
        }
    }

    // Hash Password 
    hashed, err := h.AuthService.HashPassword(req.Password)
    if err != nil {
        return nil, "", handler_models.ApiError{
            Message: "Email already exists",
            Code: http.StatusBadRequest,
        }
    } 

    user = &db_models.User{
        Name: req.Name,
        Email: req.Email,
        Password: hashed,
    }

    if err := h.UserRepo.Create(user); err != nil {
        return nil, "", handler_models.ApiError{
            Message: "Internal server error",
            Code: http.StatusInternalServerError,
        }
    }

    // Generate token
    token, err := h.AuthService.GenerateJWT(user)
    if err != nil {
        return nil, "", handler_models.ApiError{
            Message: "Error generating authentication token",
            Code: http.StatusBadRequest,
        }
    } 

    return user, token, nil 
} 
