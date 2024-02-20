package handlers

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"

	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
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

// Validate token for later

func ValidateToken(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        return []byte(env.JWT_SECRET), nil
    })
    if err != nil {
        return errors.New("Error parsing token")
    } 

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return errors.New("invalid claims")
    } 

    idFloat, ok := claims["id"].(float64)
    if !ok {
        return errors.New("invalid id claim")
    }
    id := uint(idFloat)
    if id == 0 {
        return errors.New("invalid id claim: Equals 0")
    }

    return nil

} 

func VerifyKeyFunc(t *jwt.Token) (interface{}, error) {
    return []byte(env.JWT_SECRET), nil
} 
