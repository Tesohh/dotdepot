package cmd

import (
	"fmt"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/db"
)

func Signup(userStore db.Storer[auth.User], creds auth.Credentials) error {
	err := auth.VerifyWrite(userStore, creds)
	if err != nil {
		return err
	}
	fmt.Printf("✅ signed up %v", creds.Username)
	return nil
}
