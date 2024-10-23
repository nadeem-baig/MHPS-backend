package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nadeem-baig/MHPS-backend/utils/logger"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

// getEnv retrieves the value of the environment variable named by the key or returns the default value if the variable is not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() (*sql.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	// Retrieve environment variables for PostgreSQL connection
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432") // PostgreSQL default port is 5432
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "Go_Auth")

	// Create the PostgreSQL connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	log.Printf("Attempting to connect to DB at %s:%s", host, port)

	// Open the connection to PostgreSQL
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Fatal("Failed to open database connection", err)
		return nil, err
	}

	// Ping to verify the connection
	err = db.Ping()
	if err != nil {
		logger.Fatal("Failed to ping database", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
