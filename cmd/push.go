package cmd

import (
	"fmt"

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
	fmt.Println("logged in successfully")
	return nil
}