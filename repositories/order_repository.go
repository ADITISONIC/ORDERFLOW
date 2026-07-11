package repositories

import (
	"orderflow/config"
	"orderflow/models"
)

func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}

func GetOrdersByUser(userID uint) ([]models.Order, error) {

	var orders []models.Order

	err := config.DB.
		Where("user_id = ?", userID).
		Find(&orders).Error

	return orders, err
}

func UpdateOrderStatus(orderID uint, status string) error {

	return config.DB.
		Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", status).
		Error
}