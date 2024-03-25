package middleware

import (
	"net/http"
	"strings"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authService := auth.NewAuthService()

        // Get the token from the authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return 
        } 

        // strip the "bearer" prefix from the token 
        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return 
        } 

        token := tokenParts[1]


        // Validate the token 
        err := authService.ValidateToken(token)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return 
        } 

        // Token is valid, continue to next handler
        c.Next()
    } 
} 
