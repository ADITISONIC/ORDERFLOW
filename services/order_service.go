package services

import (
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
		Quantity: quantity,
		Price: price,
		UserID: userID,
		Status: "PENDING",
	}

	return repositories.CreateOrder(&order)
}

func GetOrders(userID uint) ([]models.Order, error) {
	return repositories.GetOrdersByUser(userID)
}