package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, errPassword := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", errPassword
	}

	return string(hashPassword), nil
}

func CheckPassword(first string, second string) error {
	hashPass := []byte(first)
	bodyPass := []byte(second)
	errorHash := bcrypt.CompareHashAndPassword(hashPass, bodyPass)
	if errorHash != nil {
		return errorHash
	}

	return nil
}
