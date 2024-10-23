package config

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nadeem-baig/MHPS-backend/utils"
)

type Handler struct {
	Mux *http.ServeMux
	DB  *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		Mux: http.NewServeMux(),
		DB:  db,
	}
}

// ServeHTTP allows the handler to be used as an HTTP handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Mux.ServeHTTP(w, r)
}

// Response is a struct to define the structure of JSON responses.
type Response struct {
	Message  string      `json:"message,omitempty"`  // Pointer makes it optional, `omitempty` excludes it if nil
	Response interface{} `json:"response,omitempty"` // `interface{}` allows any type
}

type AppConfig struct {
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var AppConfigs = initAppConfigs()

func initAppConfigs() AppConfig {
	godotenv.Load()
	return AppConfig{
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              utils.GetEnv("JWT_SECRET", ""), // Use getEnv to fetch just the value.
	}
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64) // Specify base 10
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
