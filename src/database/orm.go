package database

import (
	"fmt"
	"os"

    "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const TEST_DB_PATH = "../.env.test"
const DEV_DB_PATH = "../../.env"

func ConnectDB(envPath string) (*gorm.DB, error) {
    // Load .env
    var path string
    if envPath == "test" {
        path = TEST_DB_PATH
    } else {
        path = DEV_DB_PATH
    } 


    err := godotenv.Load(path)
    if err != nil {
        return nil, fmt.Errorf("Failed to load environment variables: %w", err)
    } 

    connectionString := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s",
        os.Getenv("PG_HOST"),
        os.Getenv("PG_PORT"),
        os.Getenv("PG_USER"),
        os.Getenv("PG_PASSWORD"),
        os.Getenv("PG_DB"),
    )

    fmt.Print(connectionString)
    db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

    if err != nil {
        panic("failed to connect to database")
    } 
    return db, nil
} 
