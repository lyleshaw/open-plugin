package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password string) (string, error) {
	hashedPasswordWithSalt, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordWithSalt), nil
}

func VerifyPassword(password string, hashedPasswordWithSalt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordWithSalt), []byte(password))
	if err != nil {
		return false
	}

	return true
}
