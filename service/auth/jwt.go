package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nadeem-baig/MHPS-backend/config"
	"github.com/nadeem-baig/MHPS-backend/types"
	"github.com/nadeem-baig/MHPS-backend/utils"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.AppConfigs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func VerifyJWT(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		if tokenString == "" {
			fmt.Println("Missing or invalid JWT token")
			permissionDenied(w)
			return
		}
		token, err := validateToken(tokenString)
		if err != nil {
			fmt.Println("Invalid JWT token")
			permissionDenied(w)
			return
		}

		if !token.Valid {
			fmt.Println("Invalid JWT token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, _ := strconv.Atoi(str)
		// Convert userID to string
		userIDStr := strconv.Itoa(userID)

		// Use the converted string
		user, err := store.GetUserByID(userIDStr)
		if err != nil {
			fmt.Println("User not found")
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user)
		r = r.WithContext(ctx)
		handlerFunc(w, r)

	}

}

func getTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	return strings.Split(authHeader, " ")[1]
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.AppConfigs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.JSONResponse(w, config.Response{Message: "UnAuthorized request"}, http.StatusUnauthorized)
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
func GetUserIDFromHeaders(r *http.Request) int {
	// Assuming you have a header called "X-User-ID"
	userIDStr := r.Header.Get("X-User-ID")

	// If the header is not present or is empty
	if userIDStr == "" {
		return -1
	}

	// Convert the string header to an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		// Return -1 if the conversion fails
		return -1
	}

	return userID
}
