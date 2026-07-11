package handlers

import (
	"net/http"

	"orderflow/dto"
	"orderflow/services"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	var req dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	err := services.CreateOrder(
		req.ProductName,
		req.Quantity,
		req.Price,
		userID.(uint),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Order created successfully",
	})
}

func GetOrders(c *gin.Context) {

	userID, _ := c.Get("userID")

	orders, err := services.GetOrders(userID.(uint))

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": orders,
	})
}