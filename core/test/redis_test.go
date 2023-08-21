package test

import (
	"cloud-disk/core/db"
	"context"
	"testing"
)

func init() {
	db.InitRedisForTest("192.168.200.131", 6379)
}

func TestRedis(t *testing.T) {
	err := db.RedisClient.Set(context.Background(), "la", "la", 0).Err()
	if err != nil {
		t.Fatal(err)
	}
}
