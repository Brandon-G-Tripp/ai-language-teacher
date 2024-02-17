package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

    "github.com/Brandon-G-Tripp/ai-language-teacher/env"
)

func ConnectDB(envStr string) (*gorm.DB, error) {

    var connectionString string

    if envStr == "test" {
        connectionString = fmt.Sprintf(
            "host=%s port=%s user=%s password=%s dbname=%s",
            env.PG_HOST_TEST,
            env.PG_PORT_TEST,
            env.PG_USER_TEST,
            env.PG_PASSWORD_TEST,
            env.PG_DB_TEST,
        )
    } else {
        connectionString = fmt.Sprintf(
            "host=%s port=%s user=%s password=%s dbname=%s",
            env.PG_HOST,
            env.PG_PORT,
            env.PG_USER,
            env.PG_PASSWORD,
            env.PG_DB,
        )
    }

    fmt.Print(connectionString)
    db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

    if err != nil {
        panic("failed to connect to database")
    } 
    return db, nil
} 
