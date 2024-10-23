package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPass string, plain string) bool {
	// Compare the hashed password directly with the plain password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plain))
	return err == nil
}



