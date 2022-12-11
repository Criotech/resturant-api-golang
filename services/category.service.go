package services

import (
	"context"
	"errors"
	"github/criotech/resturant-api/models"
	"github/criotech/resturant-api/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService interface {
	CreateCategory(*models.Category) (*mongo.InsertOneResult, error)
	GetCategory(*string) (*models.Category, error)
	GetCategories() ([]*models.Category, error)
	UpdateCategory(*string, *types.UpdateCategoryRequest) error
	DeleteCategory(*string) error
}

type CategoryServiceImpl struct {
	categoryCollection *mongo.Collection
	ctx                context.Context
}

func NewCategoryServiceImpl(categoryCollection *mongo.Collection, ctx context.Context) CategoryService {
	return &CategoryServiceImpl{
		categoryCollection: categoryCollection,
		ctx:                ctx,
	}
}

func (c *CategoryServiceImpl) CreateCategory(category *models.Category) (*mongo.InsertOneResult, error) {
	category.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	category.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res, err := c.categoryCollection.InsertOne(c.ctx, category)
	return res, err
}

func (c *CategoryServiceImpl) GetCategories() ([]*models.Category, error) {
	var categories []*models.Category
	cursor, err := c.categoryCollection.Find(c.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var category models.Category
		err := cursor.Decode(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(categories) == 0 {
		return nil, errors.New("documents not found")
	}
	return categories, nil
}

func (c *CategoryServiceImpl) GetCategory(categoryID *string) (*models.Category, error) {
	var category *models.Category
	objID, err := primitive.ObjectIDFromHex(*categoryID)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objID}
	err = c.categoryCollection.FindOne(c.ctx, query).Decode(&category)
	return category, err
}

func (c *CategoryServiceImpl) UpdateCategory(categoryID *string, category *types.UpdateCategoryRequest) error {
	objID, err := primitive.ObjectIDFromHex(*categoryID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{"$set": bson.M{"name": category.Name, "updated_at": updated_at}}

	result, _ := c.categoryCollection.UpdateOne(c.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (c *CategoryServiceImpl) DeleteCategory(categoryID *string) error {
	objID, err := primitive.ObjectIDFromHex(*categoryID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, _ := c.categoryCollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
