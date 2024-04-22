package main

import (
	"log"

	"github.com/Brandon-G-Tripp/ai-language-teacher/env"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
	"github.com/gin-gonic/gin"
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

    setupRoutes(r, db)

    r.Run() // listen and serve on 0.0.0.0:8080

    sqlDB.Close()
} 
