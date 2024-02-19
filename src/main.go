package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/handlers"
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

    // Define handlers
    r.POST("api/v1/users", handlers.SignUp)

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080

    sqlDB.Close()
} 
