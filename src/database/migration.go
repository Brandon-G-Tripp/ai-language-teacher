package database

import (
	"fmt"
	"strings"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

func Migrate(envPath string) error {
    var path string
    if envPath == "test" {
        path = TEST_DB_PATH
    } else {
        path = DEV_DB_PATH
    } 

    db, err := ConnectDB(path)
    if err != nil {
        return fmt.Errorf("Failed to connect to database: %w", err)
    } 

    // Determine if this is a test env
    isTestEnv := strings.Contains(envPath, "test")

    if isTestEnv {
        // Drop table if it exists (only in test env)
        db.Migrator().DropTable(&models.User{})
    } 

    // Create table (in all envs) 
    db.Migrator().CreateTable(&models.User{})

    return nil
}
