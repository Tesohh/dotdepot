package cmd

import (
	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
	"github.com/fatih/color"
)

func Pull(userStore db.Storer[auth.User], dfStore db.Storer[db.Dotfile], cfg config.Config, creds auth.Credentials) error {
	err := auth.VerifyReadOnly(userStore, creds)
	if err != nil {
		return err
	}
	color.New(color.FgYellow, color.Italic).Println("⚠️  dotdepot is not responsible for files downloaded to your machines.")

	return nil
}
