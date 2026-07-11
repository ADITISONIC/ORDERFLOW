package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"orderflow/cache"

	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc {

	return func(c *gin.Context) {

		ip := c.ClientIP()

		key := fmt.Sprintf("rate:%s", ip)

		count, err := cache.RedisClient.Incr(context.Background(), key).Result()

		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			cache.RedisClient.Expire(context.Background(), key, time.Minute)
		}

		if count > 100 {

			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Too many requests",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}