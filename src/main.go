package main

import (
	"log"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Failed to load environment variables: %v", err)
    } 
    database.Migrate("dev")

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
} 
