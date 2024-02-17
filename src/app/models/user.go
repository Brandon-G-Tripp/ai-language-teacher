package models

import (
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

type SignUpRequest struct {
    Name string
    Email string
    Password string
} 

type SignUpResponse struct {
    User db_models.User
    Token string
} 
