package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nadeem-baig/MHPS-backend/config"
	"github.com/nadeem-baig/MHPS-backend/service/product"
	"github.com/nadeem-baig/MHPS-backend/service/user"
)

// StartServer sets up the HTTP server configurations.
func StartServer(db *sql.DB) {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	mainHandler := config.NewHandler(db)

	// Initialize service handlers
	userAuthHandler := user.NewHandler(mainHandler)
	productHandler := product.NewHandler(mainHandler)

	// Mount service handlers
	mainHandler.Mux.Handle("/api/v1/users/auth/", http.StripPrefix("/api/v1/users/auth", userAuthHandler))
	mainHandler.Mux.Handle("/api/v1/products/", http.StripPrefix("/api/v1/products", productHandler))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mainHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
