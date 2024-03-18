package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to bcrypt.GenerateFromPassword %v", err)
		return ""
	}
	return string(bytes)
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
