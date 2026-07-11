package middleware

import (
	"encoding/json"
	"net/http"

	"orderflow/cache"

	"github.com/gin-gonic/gin"
)

func IdempotencyMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		key := c.GetHeader("Idempotency-Key")

		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Idempotency-Key header required",
			})
			c.Abort()
			return
		}

		redisKey := "idem:" + key

		exists, err := cache.Exists(redisKey)

		if err == nil && exists {

			value, _ := cache.GetCache(redisKey)

			var response map[string]interface{}

			if json.Unmarshal([]byte(value), &response) == nil {

				c.JSON(http.StatusOK, response)

				c.Abort()

				return
			}
		}

		c.Set("idempotencyKey", redisKey)

		c.Next()
	}
}