package redis

import (
	"context"
	"fmt"
	
	"time"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx = context.Background()
)

func InitRedis(redisAddr string) {
	
	client = redis.NewClient(&redis.Options{
			Addr: redisAddr,
	})
}

func TryDedup(userID, itemID string) (bool, error) {

	if client == nil {
			return false, fmt.Errorf("redis client not initialized")
	}

	key := fmt.Sprintf("dedup:%s:%s", userID, itemID)
	ok, err := client.SetNX(ctx, key, "1", 10*time.Minute).Result()

	if err != nil {
			return false, err
	}

	return ok, nil
}

