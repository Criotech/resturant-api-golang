package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateCategoryRequest struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name *string            `json:"name" validate:"required,min=2,max=100"`
}
