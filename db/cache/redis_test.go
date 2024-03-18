package cache_test

import (
	"go-recommendation-system/db/cache"
	"log"
	"testing"
	"time"
)

func TestRdsSet(t *testing.T) {
	err := cache.RdsSet("aaa", "GGGG", 5*time.Second)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestRdsGet(t *testing.T) {
	s, err := cache.RdsGet("aaa")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if s != nil {
		log.Printf("aaa = %v", *s)
	} else {
		t.Errorf("aaa is nil")
	}
}
