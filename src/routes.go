package main

import (
    "github.com/gin-gonic/gin"

    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/handlers"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/middleware"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
    "gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB) {
    userRepo := repositories.NewUserRepository(db)
    conversationRepo := repositories.NewConversationRepository(db)
    messageRepo := repositories.NewMessageRepository(db)
    authService := auth.NewAuthService()

    signUpHandler := handlers.NewSignUpHandler(userRepo, authService)
    loginHandler := handlers.NewLoginHandler(userRepo, authService)
    logoutHandler := handlers.NewLogoutHandler(userRepo, authService)
    conversationHandler := handlers.NewConversationHandler(conversationRepo, authService)
    messageHandler := handlers.NewMessageHandler(messageRepo, authService)

    wrapHandler := func(handler interface{}) gin.HandlerFunc {
        return func(c *gin.Context) {
            switch h := handler.(type) {
            case func(*gin.Context):
                h(c)
            case func(*gin.Context) error: 
                if err := h(c); err != nil {
                    // handle the error if needed 
                    c.AbortWithStatus(500)
                } 
            default: 
                // handle unsupported handler types
                c.AbortWithStatus(500)
            } 
        }
    } 

    // Public routes
    publicRoutes := r.Group("/api/v1")
    {
        publicRoutes.POST("/signup", wrapHandler(signUpHandler.SignUp))
        publicRoutes.POST("/login", wrapHandler(loginHandler.Login))
    }

    // Protected Routes
    protectedRoutes := r.Group("/api/v1")
    protectedRoutes.Use(middleware.AuthMiddleware())
    {
        protectedRoutes.POST("/logout", wrapHandler(logoutHandler.Logout))

        // Conversation routes
        conversationRoutes := protectedRoutes.Group("/conversations")
        {
            conversationRoutes.POST("/", wrapHandler(conversationHandler.CreateConversation))
            conversationRoutes.GET("/:conversation_id", wrapHandler(conversationHandler.GetConversation))
            conversationRoutes.PUT("/:conversation_id", wrapHandler(conversationHandler.UpdateConversation))
            conversationRoutes.DELETE("/:conversation_id", wrapHandler(conversationHandler.DeleteConversation))
        }

        // Message routes
        messageRoutes := protectedRoutes.Group("/conversations/:conversation_id/messages")
        {
            messageRoutes.POST("/", wrapHandler(messageHandler.CreateMessage))
            messageRoutes.GET("/", wrapHandler(messageHandler.GetMessagesByConversationID))
            messageRoutes.PUT("/:message_id", wrapHandler(messageHandler.UpdateMessage))
            messageRoutes.DELETE("/:message_id", wrapHandler(messageHandler.DeleteMessage))
        }
    }

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })
}
