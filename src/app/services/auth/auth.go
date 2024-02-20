package auth

import db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"

type AuthService struct {}

func NewAuthService() *AuthService {
    return &AuthService{}
} 

func (s *AuthService) HashPassword(password string) (string, error) {
    return hashPassword(password)
} 

func (s *AuthService) ValidatePassword(hashedPassword, plainPassword string) bool {
    err := verifyPassword(hashedPassword, plainPassword)
    if err != nil {
        return false
    }
    return true
} 

func (s *AuthService) GenerateJWT(user *db_models.User) (string, error) {
    token, err := generateToken(user)   
    if err != nil {
        return "", err
    } 

    return token, nil
} 
