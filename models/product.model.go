package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CategoryID  primitive.ObjectID `json:"category_id" validate:"required"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description"`
	Price       float64            `json:"price" validate:"required"`
	Quantity    int                `json:"quantity" validate:"required"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
