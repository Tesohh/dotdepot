package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore[T any] struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (m MongoStore[T]) Get(q Query) (*T, error) {
	res := m.Coll.FindOne(context.Background(), q.ToMongo())
	var document T
	res.Decode(&document)
	// TODO: if document is empty return an error
	return &document, nil
}

func (m MongoStore[T]) Put(doc T) error {
	_, err := m.Coll.InsertOne(context.Background(), doc, nil)
	return err
}

func (m MongoStore[T]) Update(q Query, newValue T) error {
	update := bson.M{"$set": newValue}
	_, err := m.Coll.UpdateOne(context.Background(), q.ToMongo(), update)
	return err
}

func (m MongoStore[T]) Delete(q Query) error {
	_, err := m.Coll.DeleteOne(context.Background(), q.ToMongo())
	return err
}
