package main

import (
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nadeem-baig/MHPS-backend/db"
)

func main() {
	// Connect to PostgreSQL database
	dbConn, err := db.ConnectDB() // Assuming ConnectDB returns *sql.DB
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Close database connection when main function exits
	defer dbConn.Close()

	// Create a new migration driver instance for PostgreSQL
	driver, err := postgresMigrate.WithInstance(dbConn, &postgresMigrate.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	// Initialize migration with PostgreSQL instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // Path to migrations
		"postgres",                      // Database name
		driver,                          // PostgreSQL driver instance
	)
	if err != nil {
		log.Fatal("Failed to initialize migration:", err)
	}

	// Get migration version and dirty status
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal("Failed to get migration version:", err)
	}

	log.Printf("Current migration version: %d, dirty: %v", version, dirty)

	// Parse command-line arguments for "up" or "down"
	if len(os.Args) < 2 {
		log.Fatal("No migration command provided (use 'up' or 'down')")
	}

	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration failed:", err)
		} else {
			log.Println("Migration up completed successfully")
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration failed:", err)
		} else {
			log.Println("Migration down completed successfully")
		}
	default:
		log.Fatalf("Unknown command: %s (use 'up' or 'down')", cmd)
	}
}
