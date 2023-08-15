package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() (*mongo.Client, error) {
	pw := os.Getenv("mongopw")
	uri := fmt.Sprintf("mongodb+srv://tesohh:%s@cluster0.ojujhay.mongodb.net/?retryWrites=true&w=majority", pw)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, err
}
