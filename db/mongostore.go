package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrDocumentNotFound = errors.New("couldn't find document")

type IsEmptyer interface {
	IsEmpty() bool
}

type MongoStore[T IsEmptyer] struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (m MongoStore[T]) Get(q Query) (*T, error) {
	res := m.Coll.FindOne(context.Background(), q.ToMongo())
	var document T
	res.Decode(&document)

	if document.IsEmpty() {
		return nil, ErrDocumentNotFound
	}
	return &document, nil
}

func (m MongoStore[T]) GetMany(q Query) ([]T, error) {
	cur, err := m.Coll.Find(context.Background(), q.ToMongo())
	if err != nil {
		return nil, err
	}

	var results []T
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
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
