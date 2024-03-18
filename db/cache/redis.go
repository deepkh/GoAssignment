package cache

import (
	"context"
	"fmt"
	"go-recommendation-system/protos"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	RecommendationsKey = "Recommendations"
)

var rds *redis.Client
var ctx = context.Background()

func init() {
	opts := &redis.Options{
		Addr:         "127.0.0.1:6379",
		Username:     "",
		Password:     "1234", // no password set
		DB:           0,      // use default DB
		MinIdleConns: 20,
		PoolSize:     1000,
	}
	rds = redis.NewClient(opts)
}

func Rds() *redis.Client {
	return rds
}

func RdsSet(key string, val string, expr time.Duration) error {

	err := rds.Set(ctx, key, val, expr).Err()
	if err != nil {
		return fmt.Errorf("RdsSet: failed to Set: key = %v val = %v err = %v", key, val, err)
	}

	return nil
}

func RdsGet(key string) (*string, error) {
	val, err := rds.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		//log.Printf("RdsGet: key '%v' doesn't exist.", key)
		return nil, nil
	case err != nil:
		err_str := fmt.Sprintf("RdsGet: failed to Get: key = %v err = %v", key, err)
		log.Fatalf(err_str)
	}

	return &val, nil
}

func QueryRecommendations() ([]*protos.Recommendation, error) {

	// get number of records from cache
	ln, err1 := rds.LLen(context.Background(), RecommendationsKey).Result()
	if err1 != nil {
		return nil, err1
	}

	if ln == 0 {
		return nil, nil
	}

	// Get all records from cache
	ls, err4 := rds.LRange(context.Background(), RecommendationsKey, 0, -1).Result()
	if err4 != nil {
		return nil, err4
	}

	rl := make([]*protos.Recommendation, len(ls))
	for i := range ls {
		rl[i] = &protos.Recommendation{PromotionMessages: ls[i]}
	}

	return rl, nil
}

func SetRecommendations(rl []*protos.Recommendation, expr time.Duration) error {

	rds.Del(context.Background(), RecommendationsKey)

	for i := range rl {
		//log.Printf("%v = %v", i, rl[i])
		err2 := rds.LPush(context.Background(), RecommendationsKey, rl[i].PromotionMessages).Err()
		if err2 != nil {
			return err2
		}
	}

	err3 := rds.Expire(context.Background(), RecommendationsKey, expr).Err()
	if err3 != nil {
		return err3
	}

	return nil
}
