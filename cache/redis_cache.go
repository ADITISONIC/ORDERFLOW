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

	fmt.Println("Redis Connected")
}

func SetCache(key string, value string, ttl time.Duration) error {

	return RedisClient.Set(
		ctx,
		key,
		value,
		ttl,
	).Err()
}

func Exists(key string) (bool, error) {

	count, err := RedisClient.Exists(ctx, key).Result()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetCache(key string) (string, error) {

	return RedisClient.Get(ctx, key).Result()
}

func DeleteCache(key string) error {

	return RedisClient.Del(ctx, key).Err()
}