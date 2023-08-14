package main

import (
	"context"

	"github.com/Tesohh/dotdepot/cli/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client, err := db.NewMongoClient()
	if err != nil { // the universe collapses if we dont have the client
		panic(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	store := db.MongoStore{Client: client, Coll: client.Database("main").Collection("files")}
	_ = store
}
