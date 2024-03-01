package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/handlers"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/middleware"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"

	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

func main() {
    env.LoadEnv()

    db, err := database.ConnectDB("dev")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    } 

    database.Migrate("dev")

    sqlDB, err := db.DB()
    defer sqlDB.Close()

    r := gin.Default()

    userRepo := user_repo.NewUserRepository(db)
    authService := auth.NewAuthService()

    signUpHandler := handlers.NewSignUpHandler(userRepo, authService)
    loginHandler := handlers.NewLoginHandler(userRepo, authService)
    logoutHandler := handlers.NewLogoutHandler(userRepo, authService)

    // Define handlers
    r.POST("api/v1/signup", func(c *gin.Context) {
        var req handler_models.SignUpRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, err)
            return 
        } 

        user, token, err := signUpHandler.SignUp(req)
        if err != nil {
            apiErr := err.(handler_models.ApiError)
            c.JSON(apiErr.Code, apiErr.Message)
        } 

        c.JSON(http.StatusOK, handler_models.SignUpResponse{
            User: user,
            Token: token,
        })
    })

    r.POST("api/v1/login", func(c *gin.Context) {
        var req handler_models.LoginRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, err)
            return 
        } 

        user, token, err := loginHandler.Login(req.Email, req.Password)
        if err != nil {
            apiErr, ok := err.(handler_models.ApiError)
            if !ok {
                c.JSON(http.StatusInternalServerError, "Internal server error")
            } else {
                c.JSON(apiErr.Code, apiErr.Message)
            }
            return 
        } 

        c.JSON(http.StatusOK, handler_models.LoginResponse{
            User: user, 
            Token: token,
        })
    })

    r.Use(middleware.AuthMiddleware())
    r.POST("api/v1/logout", func(c *gin.Context) {
        var req handler_models.LogoutRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, err)
            return 
        } 

        err := logoutHandler.Logout(req.Token)
        if err != nil {
            apiErr, ok := err.(handler_models.ApiError)
            if !ok {
                c.JSON(http.StatusInternalServerError, "Internal server error")
            } else {
                c.JSON(apiErr.Code, apiErr.Message)
            } 
            return 
        } 
        
        c.JSON(http.StatusOK, "Successfully logged out")
    })

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080

    sqlDB.Close()
} 
