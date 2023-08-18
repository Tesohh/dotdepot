package cmd

import (
	"fmt"
	"os"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
)

func Push(userStore db.Storer[auth.User], dfStore db.Storer[db.Dotfile], cfg config.Config, creds auth.Credentials) error {
	err := auth.VerifyReadOnly(userStore, creds)
	if err != nil {
		return err
	}
	err = auth.VerifyWrite(userStore, creds)
	if err != nil {
		return err
	}
	fmt.Println("dotdepot is not responsible for files uploaded to the service.")

	filesIgnored := 0

	for _, paths := range cfg.Files {
		path, err := paths.ForCurrentOS()
		if err != nil {
			return err
		}

		if path == "" {
			filesIgnored += 1
			continue
		}

		fileBuffer, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("❌ couldn't find %v\n", path)
			continue
		}

		content := string(fileBuffer)
		df := db.Dotfile{
			UserName: creds.Username,
			Content:  content,
			Paths:    paths,
		}

		query := db.Query{"paths": paths.ToQuery(), "username": creds.Username}
		_, err = dfStore.Get(query)
		if err != nil { // this means we must create the document
			dfStore.Put(df)
		} else { // the document already exists
			dfStore.Update(query, df)
		}

		fmt.Printf("✅ pushed %v\n", path)
	}
	return nil
}
