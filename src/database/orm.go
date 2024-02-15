package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
    connectionString := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s",
        os.Getenv("PG_HOST"),
        os.Getenv("PG_PORT"),
        os.Getenv("PG_USER"),
        os.Getenv("PG_PASSWORD"),
        os.Getenv("PG_DB"),
    )
    db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

    if err != nil {
        panic("failed to connect to database")
    } 
    return db
} 
