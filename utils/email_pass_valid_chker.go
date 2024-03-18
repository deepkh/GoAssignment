package utils

import (
	"fmt"
	"regexp"
)

// CheckEmail validates the format of an email address
func CheckValidEmail(email string) bool {
	// Regular expression pattern for basic email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern into a regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	return regex.MatchString(email)
}

// CheckPassword validates the format of a password
func CheckValidPassword(pass string) error {

	// check len
	ln := len(pass)
	if !(ln >= 6 && ln <= 16) {
		return fmt.Errorf("Password '%v' is invald. !ln >=6 && ln<=16", pass)
	}

	// rules
	uppc := false
	dwnc := false
	spec := false

	// special character
	chkspeccf := func(c rune) bool {
		specc := "()[]{}<>+-*/?,.:;\"'_|~`! @#$%^&="
		for _, c2 := range specc {
			if c == c2 {
				return true
			}
		}
		return false
	}

	for _, c := range pass {

		// fit rules quick leave
		if spec && uppc && dwnc {
			return nil
		}

		// check special char
		if !spec {
			spec = chkspeccf(c)
			if spec {
				continue
			}
		}

		if !uppc {
			if c >= 'A' && c <= 'Z' {
				uppc = true
				continue
			}
		}

		if !dwnc {
			if c >= 'a' && c <= 'z' {
				dwnc = true
				continue
			}
		}
	}

	if spec && uppc && dwnc {
		return nil
	} else {
		return fmt.Errorf("Password '%v' is invald. spec %v or uppc %v or dwnc %v are not fit requiremet.",
			pass, spec, uppc, dwnc)
	}
}
