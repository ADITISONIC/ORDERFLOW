package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Redis Connected")
}

func SetCache(key string, value string) error {

	return RedisClient.Set(
		ctx,
		key,
		value,
		10*time.Minute,
	).Err()
}

func GetCache(key string) (string, error) {

	return RedisClient.Get(ctx, key).Result()
}

func DeleteCache(key string) error {

	return RedisClient.Del(ctx, key).Err()
}