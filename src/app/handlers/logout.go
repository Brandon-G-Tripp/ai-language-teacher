package handlers

import (
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

type LogoutHandler struct {
    UserRepo *user_repo.UserRepository
    AuthService *auth.AuthService
}

func NewLogoutHandler(repo *user_repo.UserRepository, authService *auth.AuthService) *LogoutHandler {
    return &LogoutHandler{UserRepo: repo, AuthService: authService}
} 

func (h *LogoutHandler) Logout(token string) error {
    err := h.AuthService.ValidateToken(token)
    if err != nil {
        return auth.ErrInvalidToken
    }

    h.AuthService.InvalidateToken(token)

    return nil
} 
