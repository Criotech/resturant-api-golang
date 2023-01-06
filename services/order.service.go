package services

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService interface {
	InitializePayment(email string, amount int64) ([]byte, error)
}

type OrderServiceImpl struct {
	orderCollection *mongo.Collection
	ctx             context.Context
}

func NewOrderServiceImpl(orderCollection *mongo.Collection, ctx context.Context) OrderService {
	return &OrderServiceImpl{
		orderCollection: orderCollection,
		ctx:             ctx,
	}
}

func (p *OrderServiceImpl) InitializePayment(email string, amount int64) ([]byte, error) {
	amount = amount * 100
	data := map[string]string{
		"email":  email,
		"amount": strconv.FormatInt(amount, 10),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.paystack.co/transaction/initialize", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer sk_test_08df2c40e350e1be93dd91a254e3193208451e72")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
