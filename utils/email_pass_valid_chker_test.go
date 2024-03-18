package utils_test

import (
	"go-recommendation-system/utils"
	"log"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	// Test cases with valid email addresses
	validEmails := []string{
		"test@example.com",
		"user123@email.co.uk",
		"john.doe123@example-domain.com",
		"a@b.cc",
	}

	// Test cases with invalid email addresses
	invalidEmails := []string{
		"not_an_email",
		"user@domain",
		"user@example.",
		"@example.com",
		"a@b.c",
		"a@.cc",
	}

	// Test valid email addresses
	for _, email := range validEmails {
		if !utils.CheckValidEmail(email) {
			t.Errorf("IsValidEmail(%q) = false, expected true", email)
		}
	}

	// Test invalid email addresses
	for _, email := range invalidEmails {
		if utils.CheckValidEmail(email) {
			t.Errorf("IsValidEmail(%q) = true, expected false", email)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	// Test cases with valid passwords
	validPasswords := []string{
		"123456Aa!",
		"@b012Gz12f",
		"b012\"Gz12f",
		"b012Gz12f=",
	}

	// Test cases with invalid passwords
	invalidPasswords := []string{
		"12345",
		"123456",
		"123456A",
		"123456Ab",
		"123456AbC",
		"123456abc(",
		"123456abc(d=",
	}

	// Test valid passwords
	for _, password := range validPasswords {
		e := utils.CheckPassword(password)

		if e != nil {
			log.Printf("%v", e)
			t.Errorf("CheckPassword(%q) = false, expected true err = %v", password, e)
		}
	}

	// Test invalid passwords
	for _, password := range invalidPasswords {
		e := utils.CheckPassword(password)
		if e != nil {
			log.Printf("%v", e)
		}

		if e == nil {
			t.Errorf("CheckPassword(%q) = true, expected false err = %v", password, e)
		}
	}
}
