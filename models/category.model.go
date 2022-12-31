package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category_Url struct {
	Url string `json:"url" validate:"required"`
	ID  string `json:"id" validate:"required"`
}

type Category struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         *string            `json:"name" validate:"required,min=2,max=100"`
	Category_Url Category_Url       `json:"category_url" validate:"required"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
}
