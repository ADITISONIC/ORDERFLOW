package handlers

import (
	"net/http"

	"orderflow/services"

	"github.com/gin-gonic/gin"
)

func GetMetrics(c *gin.Context) {

	metrics, err := services.GetMetrics()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": metrics,
	})
}