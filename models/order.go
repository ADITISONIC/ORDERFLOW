package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	ProductName string  `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`

	Status string `gorm:"default:'PENDING'"`

	UserID uint
	User   User
}