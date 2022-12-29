package services

import (
	"context"
	"errors"
	"fmt"
	"github/criotech/resturant-api/models"
	"github/criotech/resturant-api/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductService interface {
	CreateProduct(*models.Product) (*mongo.InsertOneResult, error)
	GetProducts(types.QueryProducts) ([]*models.Product, error)
	GetProduct(*string) (*models.Product, error)
	UpdateProduct(*string, *types.UpdateProductRequest) error
	DeleteProduct(*string) error
}

type ProductServiceImpl struct {
	productCollection *mongo.Collection
	ctx               context.Context
}

func NewProductServiceImpl(productCollection *mongo.Collection, ctx context.Context) ProductService {
	return &ProductServiceImpl{
		productCollection: productCollection,
		ctx:               ctx,
	}
}

func (p *ProductServiceImpl) CreateProduct(product *models.Product) (*mongo.InsertOneResult, error) {
	product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res, err := p.productCollection.InsertOne(p.ctx, product)
	return res, err
}

func (p *ProductServiceImpl) GetProducts(queries types.QueryProducts) ([]*models.Product, error) {
	var products []*models.Product
	fmt.Println(queries)

	filter := bson.M{}

	if queries.Category != "" {
		categoryID, err := primitive.ObjectIDFromHex(queries.Category)
		if err != nil {
			return nil, err
		}

		filter["category_id"] = categoryID
	}

	if queries.MinPrice > 0 && queries.MaxPrice > 0 {
		filter["price"] = bson.M{"$gte": queries.MinPrice, "$lte": queries.MaxPrice}
	} else if queries.MinPrice > 0 {
		filter["price"] = bson.M{"$gte": queries.MinPrice}
	} else if queries.MaxPrice > 0 {
		filter["price"] = bson.M{"$lte": queries.MaxPrice}
	}

	cursor, err := p.productCollection.Find(p.ctx, filter, options.Find().SetSkip(queries.Page).SetLimit(queries.Limit))
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(p.ctx)

	if len(products) == 0 {
		return nil, errors.New("documents not found")
	}
	return products, nil
}

func (p *ProductServiceImpl) GetProduct(productID *string) (*models.Product, error) {
	var product *models.Product
	objID, err := primitive.ObjectIDFromHex(*productID)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objID}
	err = p.productCollection.FindOne(p.ctx, query).Decode(&product)
	return product, err
}

func (p *ProductServiceImpl) UpdateProduct(productID *string, product *types.UpdateProductRequest) error {
	objID, err := primitive.ObjectIDFromHex(*productID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{"$set": bson.M{"name": product.Name, "description": product.Description, "price": product.Price, "quantity": product.Quantity, "updated_at": updated_at}}

	result, _ := p.productCollection.UpdateOne(p.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (p *ProductServiceImpl) DeleteProduct(productID *string) error {
	objID, err := primitive.ObjectIDFromHex(*productID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, _ := p.productCollection.DeleteOne(p.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
