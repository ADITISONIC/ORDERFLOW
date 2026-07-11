package services

import (
	"encoding/json"
	"fmt"
	"time"

	"orderflow/cache"
	"orderflow/events"
	"orderflow/kafka"
	"orderflow/models"
	"orderflow/repositories"
)

func CreateOrder(
	product string,
	quantity int,
	price float64,
	userID uint,
) (*models.Order, error) {

	order := models.Order{
		ProductName: product,
		Quantity:    quantity,
		Price:       price,
		UserID:      userID,
		Status:      "PENDING",
	}

	// Save order to MySQL
	if err := repositories.CreateOrder(&order); err != nil {
		return nil, err
	}

	// Publish Kafka event
	event := events.OrderCreatedEvent{
		OrderID:     order.ID,
		UserID:      order.UserID,
		ProductName: order.ProductName,
		Quantity:    order.Quantity,
		Price:       order.Price,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	if err := kafka.Publish(data); err != nil {
		return nil, err
	}

	// Invalidate cached order list
	cacheKey := fmt.Sprintf("orders:%d", userID)
	cache.DeleteCache(cacheKey)

	return &order, nil
}

func GetOrders(userID uint) ([]models.Order, error) {

	cacheKey := fmt.Sprintf("orders:%d", userID)

	// Check Redis first
	cachedData, err := cache.GetCache(cacheKey)

	if err == nil {

		var orders []models.Order

		if err := json.Unmarshal([]byte(cachedData), &orders); err == nil {
			fmt.Println("Cache Hit")
			return orders, nil
		}
	}

	fmt.Println("Cache Miss")

	// Fetch from MySQL
	orders, err := repositories.GetOrdersByUser(userID)

	if err != nil {
		return nil, err
	}

	// Cache result for 10 minutes
	data, _ := json.Marshal(orders)

	cache.SetCache(
		cacheKey,
		string(data),
		10*time.Minute,
	)

	return orders, nil
}

func UpdateStatus(orderID uint, status string) error {
	return repositories.UpdateOrderStatus(orderID, status)
}