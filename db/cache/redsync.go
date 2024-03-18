package cache

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

var rdSync *redsync.Redsync
var recomMut *redsync.Mutex

func init() {
	opts := &redis.Options{
		Addr:         "127.0.0.1:6379",
		Username:     "",
		Password:     "1234", // no password set
		DB:           0,      // use default DB
		MinIdleConns: 20,
		PoolSize:     1000,
	}
	redisConn := goredislib.NewClient(opts)

	// Create an instance of redisync to be used to obtain a mutual exclusion  lock.
	rdSync = redsync.New(goredis.NewPool(redisConn))

	// Mutex for recommendation
	recomMut = rdSync.NewMutex("recommendation",
		redsync.WithExpiry(time.Duration(9999)*time.Second),
		redsync.WithTries(9999))
}

func RecomMutLock() error {
	return recomMut.Lock()
}

func RecomMutUnLock() error {
	_, err := recomMut.Unlock()
	return err
}
