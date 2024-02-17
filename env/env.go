package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
    "github.com/Brandon-G-Tripp/ai-language-teacher/internal/projectpath"
)

var (
    // Define package-level variables for environment variables
    PG_HOST     string
    PG_PORT     string
    PG_USER     string
    PG_PASSWORD string
    PG_DB       string

    // test envs
    PG_HOST_TEST     string
    PG_PORT_TEST     string
    PG_USER_TEST     string
    PG_PASSWORD_TEST string
    PG_DB_TEST       string
)

func LoadEnv() {
    // Load .env file
    err := godotenv.Load(projectpath.Root + "/.env")
    if err != nil {
        log.Fatalf("Failed to load environment variables in env.go: %v", err)
    }

    // Assign environment variables to package-level variables
    PG_HOST = os.Getenv("PG_HOST")
    PG_PORT = os.Getenv("PG_PORT")
    PG_USER = os.Getenv("PG_USER")
    PG_PASSWORD = os.Getenv("PG_PASSWORD")
    PG_DB = os.Getenv("PG_DB")

    PG_HOST_TEST = os.Getenv("PG_HOST")
    PG_PORT_TEST = os.Getenv("PG_PORT")
    PG_USER_TEST = os.Getenv("PG_USER")
    PG_PASSWORD_TEST = os.Getenv("PG_PASSWORD")
    PG_DB_TEST = os.Getenv("PG_DB")
}
