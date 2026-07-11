package events

type OrderCreatedEvent struct {
	OrderID     uint    `json:"order_id"`
	UserID      uint    `json:"user_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}