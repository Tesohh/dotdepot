package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Query map[string]any

func (q Query) ToMongo() primitive.D {
	d := primitive.D{}
	for k, v := range q {
		d = append(d, primitive.E{Key: k, Value: v})
	}
	return d
}

func (q Query) ToParameters() map[string]string {
	m := make(map[string]string)
	for k, v := range q {
		m[k] = fmt.Sprint(v)
	}
	return m
}

type Storer[T any] interface {
	Get(Query) (*T, error)
	GetMany(Query) ([]T, error)
	Put(T) error
	Update(Query, T) error
	Delete(Query) error
}
