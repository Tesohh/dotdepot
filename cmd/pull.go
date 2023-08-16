package cmd

import (
	"fmt"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
)

func Pull(userStore db.Storer[auth.User], dfStore db.Storer[db.Dotfile], cfg config.Config) error {
	fmt.Println("pull")
	return nil
}
