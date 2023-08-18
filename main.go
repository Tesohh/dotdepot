package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/cmd"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client, err := db.NewMongoClient()
	if err != nil { // the universe collapses if we dont have the client
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	dfStore := db.MongoStore[db.Dotfile]{Client: client, Coll: client.Database("main").Collection("files")}
	userStore := db.MongoStore[auth.User]{Client: client, Coll: client.Database("main").Collection("users")}

	cfg, err := config.Read[config.Config]("config.yml")
	if err != nil {
		log.Fatal(config.ErrNoConfigFile)
	}
	creds, err := config.Read[auth.Credentials]("login.yml")
	if err != nil {
		log.Fatal(auth.ErrNoLoginFile)
	}

	fmt.Println("📦 dotdepot")
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("	- dotdepot push")
		fmt.Println("	- dotdepot pull")
		return
	}
	switch os.Args[1] {
	case "push":
		err = cmd.Push(userStore, dfStore, *cfg, *creds)
	case "pull":
		err = cmd.Pull(userStore, dfStore, *cfg)
	}
	if err != nil {
		log.Fatal(err)
	}
}
