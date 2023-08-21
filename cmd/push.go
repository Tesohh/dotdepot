package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
	"github.com/Tesohh/dotdepot/cli/walk"
	"github.com/fatih/color"
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
	color.New(color.FgYellow, color.Italic).Println("⚠️  dotdepot is not responsible for files uploaded to the service.")

	filesIgnored := 0

	// make sure the config gets pushed too
	files := cfg.Files
	cfgpath := "~/.config/dotdepot/config.yml"
	files = append(files, db.Paths{
		Windows: cfgpath,
		MacOS:   cfgpath,
		Linux:   cfgpath,
	})
	for _, dir := range cfg.Directories {
		ospath, err := dir.ForCurrentOS()
		if err != nil {
			return err
		}
		paths, err := walk.Walk(ospath, true)
		if err != nil {
			return err
		}
		for _, v := range paths {
			p := db.Paths{}
			if dir.Linux != "" {
				p.Linux = path.Join(dir.Linux, v)
			}
			if dir.MacOS != "" {
				p.MacOS = path.Join(dir.MacOS, v)
			}
			if dir.Windows != "" {
				p.Windows = path.Join(dir.Windows, v)
			}
			files = append(files, p)
		}
	}

	for _, paths := range files {
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

		// upsert
		query := db.Query{"paths": paths.ToQuery(), "username": creds.Username}
		_, err = dfStore.Get(query)
		if err != nil { // this means we must create the document
			dfStore.Put(df)
		} else { // the document already exists
			dfStore.Update(query, df)
		}

		fmt.Printf("✅ pushed %v\n", path)
	}
	if filesIgnored > 0 {
		fmt.Printf("🤷 ignored %v files", filesIgnored)
	}
	return nil
}
