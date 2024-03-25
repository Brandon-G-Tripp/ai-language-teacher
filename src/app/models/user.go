package models

import (
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

type SignUpRequest struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
} 

type SignUpResponse struct {
    User *db_models.User `json:"user"`
    Token string `json:"token"`
} 

type LoginRequest struct {
    Email string `json:"email"`
    Password string `json:"password"`
} 

type LoginResponse struct {
    User *db_models.User `json:"user"`
    Token string `json:"token"`
} 

type LogoutRequest struct {
    Token string `json:"token"`
} 

type LogoutResponse struct {
    Message string `json:"message"`
}
