package handlers

import (
	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *db_models.User) (string, error) {
    secretKey := []byte(env.JWT_SECRET)

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    // Set token claims 
    claims["id"] = user.ID
    claims["email"] = user.Email

    return token.SignedString(secretKey)
} 
