package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderedItem struct {
	ProductId   primitive.ObjectID `json:"product_id" validate:"required"`
	ProductName string             `json:"product_name" validate:"required"`
	Quantity    int64              `json:"quantity" validate:"required"`
	UnitPrice   int64              `json:"unit_price" validate:"required"`
	Price       int64              `json:"price" validate:"required"`
}

type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CustomerId   primitive.ObjectID `json:"customer_id" validate:"required"`
	OrderedItems []OrderedItem      `json:"ordered_items"`
	TotalPrice   int64              `json:"total_price" validate:"required"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
}
