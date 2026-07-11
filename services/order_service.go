package services

import (
	"encoding/json"
	"fmt"

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
) error {

	order := models.Order{
		ProductName: product,
		Quantity:    quantity,
		Price:       price,
		UserID:      userID,
		Status:      "PENDING",
	}
	if err := repositories.CreateOrder(&order); err != nil {
		return err
	}
	event := events.OrderCreatedEvent{
		OrderID:     order.ID,
		UserID:      order.UserID,
		ProductName: order.ProductName,
		Quantity:    order.Quantity,
		Price:       order.Price,
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := kafka.Publish(data); err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("orders:%d", userID)

	cache.DeleteCache(cacheKey)

	return nil
}

func GetOrders(userID uint) ([]models.Order, error) {

	cacheKey := fmt.Sprintf("orders:%d", userID)
	cachedData, err := cache.GetCache(cacheKey)

	if err == nil {

		var orders []models.Order

		err = json.Unmarshal([]byte(cachedData), &orders)

		if err == nil {
			fmt.Println("Cache Hit")
			return orders, nil
		}
	}

	fmt.Println("Cache Miss")
	orders, err := repositories.GetOrdersByUser(userID)

	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(orders)

	cache.SetCache(cacheKey, string(data))

	return orders, nil
}

func UpdateStatus(orderID uint, status string) error {
	return repositories.UpdateOrderStatus(orderID, status)
}
