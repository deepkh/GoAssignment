package db_test

import (
	"go-recommendation-system/db"
	"go-recommendation-system/protos"
	"log"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {

	err := db.CreateUser(&protos.User{Email: "hello@world.com", PasswordHashed: "passsword", Confirm: 0, Timestamp: time.Now().UnixMicro()})
	if err != nil {
		t.Errorf("%v", err)
		return
	}

}

func TestQueryUser(t *testing.T) {

	u, err := db.QueryUser("hello@world.com")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	log.Printf("User = %v", u)
}

func TestQueryRecommendations(t *testing.T) {
	_, err := db.QueryRecommendations()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
}
