package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"orderflow/cache"
	"orderflow/dto"
	"orderflow/services"

	"github.com/gin-gonic/gin"
)
// CreateOrder godoc
//
// @Summary Create Order
// @Description Creates a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateOrderRequest true "Order Details"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /orders [post]
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

	order,err := services.CreateOrder(
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

	response := gin.H{
		"success": true,
		"message": "Order created successfully",
		"orderId": order.ID,
		"status":  order.Status,
	}
	key, exists := c.Get("idempotencyKey")

	if exists {

		responseJSON, _ := json.Marshal(response)

		cache.SetCache(
			key.(string),
			string(responseJSON),
			24*time.Hour,
		)
	}
	c.JSON(http.StatusCreated, response)
}

// GetOrders godoc
//
// @Summary Get Orders
// @Description Returns all orders for the logged-in user
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /orders [get]
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
		"data":    orders,
	})
}
