package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (m MongoStore) Get(q Query) (*Dotfile, error) {
	res := m.Coll.FindOne(context.Background(), q.ToMongo())
	df := Dotfile{}
	res.Decode(&df)
	if (df == Dotfile{}) {
		return nil, fmt.Errorf("no documents exist with query %v", q)
	}
	return &df, nil
}

func (m MongoStore) Put(df Dotfile) error {
	_, err := m.Coll.InsertOne(context.Background(), df, nil)
	return err
}

func (m MongoStore) Update(q Query, newValue Dotfile) error {
	update := bson.M{"$set": newValue}
	_, err := m.Coll.UpdateOne(context.Background(), q.ToMongo(), update)
	return err
}

func (m MongoStore) Delete(q Query) error {
	_, err := m.Coll.DeleteOne(context.Background(), q.ToMongo())
	return err
}

// TODO: make Update and Delete use a Query like in Get.
