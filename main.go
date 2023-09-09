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
	"github.com/fatih/color"
	_ "github.com/joho/godotenv/autoload"
)

var endpoint = func() string {
	postfix := "/.netlify/functions"
	envPrefix := os.Getenv("endpointprefix")
	if envPrefix != "" {
		return envPrefix + postfix
	}

	return "https://dotdepot.pyros.dev" + postfix
}()

func main() {
	client, err := db.NewMongoClient()
	if err != nil { // the universe collapses if we dont have the client
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Read[config.Config]("config.yml")
	if err != nil {
		log.Fatal(config.ErrNoConfigFile)
	}
	creds, err := config.Read[auth.Credentials]("login.yml")
	if err != nil {
		log.Fatal(auth.ErrNoLoginFile)
	}
	dfStore := db.CRUDStore[db.Dotfile]{Endpoint: endpoint, Collection: "files", Username: creds.Username, Password: creds.Password}
	userStore := db.CRUDStore[auth.User]{Endpoint: endpoint, Collection: "users", Username: creds.Username, Password: creds.Password}

	color.New(color.Bold).Println("📦 dotdepot")
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
		err = cmd.Pull(userStore, dfStore, *creds)
	case "signup":
		err = cmd.Signup(userStore, *creds)
	}
	if err != nil {
		log.Fatal(err)
	}
}
