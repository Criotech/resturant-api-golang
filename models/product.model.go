package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	CategoryID  primitive.ObjectID `json:"category_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Quantity    int                `json:"quantity"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
