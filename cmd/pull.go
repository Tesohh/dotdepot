package cmd

import (
	"fmt"
	"os"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

func Pull(userStore db.Storer[auth.User], dfStore db.Storer[db.Dotfile], creds auth.Credentials) error {
	err := auth.VerifyReadOnly(userStore, creds)
	if err != nil {
		return err
	}
	color.New(color.FgYellow, color.Italic).Println("⚠️  dotdepot is not responsible for files downloaded to your machines.")

	// first of all we need the new config:
	cfgpath := "~/.config/dotdepot/config.yml"
	query := db.Query{
		"paths": db.Paths{
			Windows: cfgpath,
			MacOS:   cfgpath,
			Linux:   cfgpath,
		}.ToQuery(),
		"username": creds.Username,
	}
	rawCfg, err := dfStore.Get(query)
	if err != nil {
		return err
	}
	var cfg config.Config
	yaml.Unmarshal([]byte(rawCfg.Content), &cfg)

	filesIgnored := 0
	for _, paths := range cfg.Files {
		query := db.Query{"paths": paths.ToQuery(), "username": creds.Username}
		df, err := dfStore.Get(query)
		if err != nil {
			return err
		}

		path, err := df.Paths.ForCurrentOS()
		if err != nil {
			return err
		}

		if path == "" {
			filesIgnored += 1
			continue
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(df.Content)
		if err != nil {
			return err
		}

		fmt.Printf("✅ pulled %v\n", path)
	}

	if filesIgnored > 0 {
		fmt.Printf("🤷 ignored %v files", filesIgnored)
	}

	return nil
}
