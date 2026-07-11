package repositories

import (
	"orderflow/config"
	"orderflow/models"
)

func CountOrders() (int64, error) {

	var count int64

	err := config.DB.
		Model(&models.Order{}).
		Count(&count).Error

	return count, err
}

func CountOrdersByStatus(status string) (int64, error) {

	var count int64

	err := config.DB.
		Model(&models.Order{}).
		Where("status = ?", status).
		Count(&count).Error

	return count, err
}