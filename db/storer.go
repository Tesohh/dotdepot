package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Query map[string]any

func (q Query) ToMongo() primitive.D {
	d := primitive.D{}
	for k, v := range q {
		d = append(d, primitive.E{Key: k, Value: v})
	}
	return d
}

type Storer[T any] interface {
	Get(Query) (*T, error)
	Put(T) error
	Update(Query, T) error
	Delete(Query) error
}
