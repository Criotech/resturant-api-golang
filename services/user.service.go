package services

import (
	"context"
	"errors"
	"fmt"
	"github/criotech/resturant-api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateAccount(*models.User) (*mongo.InsertOneResult, error)
	GetUserByEmail(*string) (*models.User, error)
	UpdateUserTokens(string, string, primitive.ObjectID) error
}

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserServiceImpl(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (c *UserServiceImpl) CreateAccount(product *models.User) (*mongo.InsertOneResult, error) {
	product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res, err := c.userCollection.InsertOne(c.ctx, product)
	return res, err
}

func (p *UserServiceImpl) GetUserByEmail(email *string) (*models.User, error) {
	var user *models.User

	query := bson.M{"email": *email}
	err := p.userCollection.FindOne(p.ctx, query).Decode(&user)

	return user, err
}

func (p *UserServiceImpl) UpdateUserTokens(signedToken string, signedRefreshToken string, userId primitive.ObjectID) error {
	fmt.Println(userId, "this is my id")

	filter := bson.M{"_id": userId}
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{"$set": bson.M{"token": signedToken, "refresh_token": signedRefreshToken, "updated_at": updated_at}}

	result, _ := p.userCollection.UpdateOne(p.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
