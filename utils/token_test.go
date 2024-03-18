package utils_test

import (
	"log"
	"testing"
	"time"

	"go-recommendation-system/utils"
)

const testEmail = "Hello@world.com"

func TestGenerateTokenAndVerify(t *testing.T) {

	// generate
	token, err := utils.GenereUserToken(testEmail, time.Second*5)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	log.Printf("Token = %v", token)

	// verify & decrypt
	email, err1 := utils.ParseUserToken(token)
	if err1 != nil {
		t.Errorf("%v", err1)
		return
	}

	log.Printf("verify & decrypt, email = %v", email)

	// wait until expired
	log.Printf("wait 5 seconds until expired ...")
	time.Sleep(time.Second * 5)

	// verify & decrypt
	email2, err2 := utils.ParseUserToken(token)
	if err2 == nil {
		t.Errorf("should be expired. but not. email2 = %v", email2)
		return
	}

	log.Printf("Verify again. expected it to have expired. = '%v'", err2)

}
