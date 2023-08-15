package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Tesohh/dotdepot/cli/auth"
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

	// dfstore := db.MongoStore{Client: client, Coll: client.Database("main").Collection("files")}
	userStore := db.MongoStore[auth.User]{Client: client, Coll: client.Database("main").Collection("users")}
	creds, err := auth.Read()
	if err != nil {
		log.Fatal("Please create a login.yml in the dotdepot folder (~/.config/dotdepot)")
	}
	if creds.Username == "" {
		log.Fatal("Please specify at least a username in your login.yml to get read access.")
	}

	err = auth.Verify(userStore, creds)
	fmt.Println(err)
}
