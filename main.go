package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
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
	creds, err := config.Read[auth.Credentials]("login.yml")

	err = auth.VerifyReadOnly(userStore, *creds, err)
	if err != nil {
		log.Fatal(err)
	}
	err = auth.VerifyWrite(userStore, *creds)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("You have read access to %v, but need write access to do this action.", creds.Username)
	}
	fmt.Println("Logged in! Write access")
}
