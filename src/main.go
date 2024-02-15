package main

import "github.com/gin-gonic/gin"
import "github.com/Brandon-G-Tripp/ai-language-teacher/src/database"

func main() {
    database.AutoMigrate()

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })

    r.Run() // listen and serve on 0.0.0.0:8080
} 
