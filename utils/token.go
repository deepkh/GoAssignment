package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("f7e2f8cb-d5a6-495d-bece-47e100455ac1")

type UserTokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenereUserToken(email string, expr time.Duration) (t string, err error) {
	// Create claims with multiple fields populated
	utc := UserTokenClaims{
		email,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expr)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, utc)
	t, err = token.SignedString(mySigningKey)
	if err != nil {
		return
	}

	return
}

func ParseUserToken(t string) (email string, err error) {

	token, err := jwt.ParseWithClaims(t, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return
	} else if claims, ok := token.Claims.(*UserTokenClaims); ok {
		email = claims.Email
		return
	} else {
		err = fmt.Errorf("ParseUserToken: unknown claims type, cannot proceed: %v", t)
	}

	return
}
