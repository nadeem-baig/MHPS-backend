package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/nadeem-baig/MHPS-backend/utils/logger"
)

// JSONResponse sends a JSON response with the given status code.
func JSONResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return logger.Errorf("Missing Body in the request")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

var Validate = validator.New()

// Helper function to get environment variable as a string with fallback
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
