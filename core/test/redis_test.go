package test

import (
	"cloud-disk/core/db"
	"context"
	"testing"
)

func init() {
	db.InitRedisForTest("localhost", 6379)
}

func TestRedis(t *testing.T) {
	err := db.RedisClient.Set(context.Background(), "la", "la", 0).Err()
	if err != nil {
		t.Fatal(err)
	}
}
