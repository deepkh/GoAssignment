package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-recommendation-system/db/cache"
	"go-recommendation-system/protos"
)

var db *gorm.DB = nil

func init() {
	var err error
	db, err = gorm.Open(
		mysql.Open("root:1234@tcp(127.0.0.1:3306)/RecommendationSystem?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		log.Panicf("%v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("%v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func CreateUser(u *protos.User) error {
	r := db.Create(u)

	if r.Error != nil {
		return r.Error
	}

	return nil
}

func QueryUser(email string) (*protos.User, error) {
	u := protos.User{}
	r := db.First(&u, "email = ?", email)

	if r.Error != nil {
		// It's allowing the user to be nil.
		if r.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, r.Error
	}

	// This case should not have happened.
	if r.RowsAffected > 1 {
		return nil, fmt.Errorf("failed to find %v in the db. r.RowsAffected %d", email, r.RowsAffected)
	}

	return &u, nil
}

func UpdateUser(u *protos.User) (err error) {

	r := db.Model(&protos.User{}).Where("email = ?", u.Email).Updates(u)
	if r.Error != nil {
		err = r.Error
		return
	}

	// This case should not have happened.
	if r.RowsAffected > 1 {
		err = fmt.Errorf("failed to update %v in the db. r.RowsAffected %d", u, r.RowsAffected)
		return
	}

	return
}

func QueryRecommendations() ([]*protos.Recommendation, error) {

	// TODO Use the single locker will down performance dramatically.
	// TODO It would be better use read/write locker instead of single locker.
	err := cache.RecomMutLock()
	if err != nil {
		return nil, err
	}

	defer func() {
		err := cache.RecomMutUnLock()
		if err != nil {
			log.Printf("failed to RecomMutUnLock %v", err)
			return
		}
	}()

	// check from cache
	rl, err := cache.QueryRecommendations()
	if err != nil {
		return nil, err
	}

	// use cache
	if rl != nil {
		log.Printf("Get %v of Recommendations from cache", len(rl))
		return rl, nil
	}

	// use db
	rl2 := []*protos.Recommendation{}
	r := db.Find(&rl2)
	if r.Error != nil {
		return nil, r.Error
	}

	// update to cache as well
	expr := 600
	err3 := cache.SetRecommendations(rl2, time.Duration(expr)*time.Second)
	if err3 != nil {
		return nil, err3
	}

	log.Printf("\t\t\t Set %v of Recommendations to cache and set expr to %v seconds", len(rl2), expr)
	return rl2, nil
}
