package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type MongoStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (m MongoStore) Get(filename string, username string) Dotfile {
	res := m.Coll.FindOne(context.Background(), bson.D{{Key: "filename", Value: filename}, {Key: "username", Value: username}})
	df := Dotfile{}
	res.Decode(&df)
	return df
}

func (m MongoStore) Put(df Dotfile) error {
	_, err := m.Coll.InsertOne(context.Background(), df, nil)
	return err
}

func (m MongoStore) Update(id primitive.ObjectID, newValue Dotfile) error {
	update := bson.M{"$set": newValue}
	_, err := m.Coll.UpdateByID(context.Background(), id, update, nil)
	return err
}

func (m MongoStore) Delete(id primitive.ObjectID) error {
	_, err := m.Coll.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}

type DFStorer interface {
	Get(string, string) Dotfile
	Put(Dotfile) error
	Update(primitive.ObjectID, Dotfile) error
	Delete(primitive.ObjectID) error
}
