package cache_test

import (
	"context"
	"go-recommendation-system/db/cache"
	"log"
	"sync"
	"testing"
)

func TestRecomMutLock(t *testing.T) {

	// Lock func
	lckf := func(i int) {
		err := cache.RecomMutLock()
		if err != nil {
			t.Errorf("lckf i = %v, err = %v", i, err)
			return
		}
	}

	// Ublock func
	unlckf := func(i int) {
		err := cache.RecomMutUnLock()
		if err != nil {
			t.Errorf("unlckf i = %v, err = %v", i, err)
			return
		}
	}

	// number of producers
	numProducers := 10000
	wg := new(sync.WaitGroup)
	wg.Add(numProducers)

	cache.Rds().Del(context.Background(), "SUM")

	// producer function
	prdf := func(i int) {
		defer wg.Done()
		lckf(i)
		defer unlckf(i)
		_, err := cache.Rds().IncrBy(context.Background(), "SUM", int64(i)).Result()
		if err != nil {
			t.Errorf("failed to cache.Rds().Incr %v", err)
		}
	}

	for i := range numProducers {
		go prdf(i + 1)
	}

	wg.Wait()
	sum3, err := cache.Rds().Get(context.Background(), "SUM").Int64()
	if err != nil {
		t.Errorf("failed to Rds().Get %v", err)
	}

	var sum2 int64 = (1 + int64(numProducers)) * int64(numProducers) / 2
	if sum3 == sum2 {
		log.Printf("sum %v == %v", sum3, sum2)
	} else {
		t.Fatalf("sum %v != %v", sum3, sum2)
	}
}
