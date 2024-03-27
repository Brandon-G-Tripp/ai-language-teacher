package database

import (
	"fmt"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

func Migrate(env string) error {

    db, err := ConnectDB(env)
    if err != nil {
        return fmt.Errorf("Failed to connect to database: %w", err)
    } 


    if env == "test" {
        // Drop table if it exists (only in test env)
        db.Migrator().DropTable(&models.User{})
        db.Migrator().DropTable(&models.Conversation{})
        db.Migrator().DropTable(&models.Message{})
    } 

    // Create table (in all envs) 
    db.Migrator().CreateTable(&models.User{})
    db.Migrator().CreateTable(&models.Conversation{})
    db.Migrator().CreateTable(&models.Message{})

    return nil
}
