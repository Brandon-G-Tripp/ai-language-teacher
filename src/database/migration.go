package database

import (
	"fmt"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

func AutoMigrate() error {
    db, err := ConnectDB("../../.env")
    if err != nil {
        return fmt.Errorf("Failed to connect to database: %w", err)
    } 
    db.AutoMigrate(&models.User{})
    return nil
}
