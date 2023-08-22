package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/Tesohh/dotdepot/cli/auth"
	"github.com/Tesohh/dotdepot/cli/config"
	"github.com/Tesohh/dotdepot/cli/db"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

// Absolutely disgusting implementation but it works i guess
func getFilesFromDirs(dfStore db.Storer[db.Dotfile], cfg config.Config, creds auth.Credentials) ([]db.Paths, error) {
	dfs := make([]db.Paths, 0)
	for _, dir := range cfg.Directories {
		res, err := dfStore.GetMany(db.Query{"username": creds.Username})
		if err != nil {
			return nil, err
		}

		for _, r := range res {
			rp, err := r.Paths.ForCurrentOS()
			if err != nil {
				return nil, err
			}
			dp, err := dir.ForCurrentOS()
			if err != nil {
				return nil, err
			}
			if rp == "" || dp == "" || rp == "." || dp == "." {
				continue
			}
			if strings.Contains(rp, dp) {
				dfs = append(dfs, r.Paths)
			}
		}
	}
	return dfs, nil
}

func Pull(userStore db.Storer[auth.User], dfStore db.Storer[db.Dotfile], creds auth.Credentials) error {
	err := auth.VerifyReadOnly(userStore, creds)
	if err != nil {
		return err
	}
	color.New(color.FgYellow, color.Italic).Println("⚠️  dotdepot is not responsible for files downloaded to your machines.")

	// first of all we need the new config:
	cfgpath := "~/.config/dotdepot/config.yml"
	cfgpaths := db.Paths{
		Windows: cfgpath,
		MacOS:   cfgpath,
		Linux:   cfgpath,
	}
	query := db.Query{
		"paths":    cfgpaths.ToQuery(),
		"username": creds.Username,
	}
	rawCfg, err := dfStore.Get(query)
	if err != nil {
		return err
	}
	var cfg config.Config
	err = yaml.Unmarshal([]byte(rawCfg.Content), &cfg)
	if err != nil {
		return err
	}

	files := cfg.Files
	files = append(files, cfgpaths)

	res, err := getFilesFromDirs(dfStore, cfg, creds)
	if err != nil {
		return err
	}
	files = append(files, res...)

	filesIgnored := 0
	for _, paths := range files {
		query := db.Query{"paths": paths.ToQuery(), "username": creds.Username}
		df, err := dfStore.Get(query)
		if err != nil {
			return err
		}

		p, err := df.Paths.ForCurrentOS()
		if err != nil {
			return err
		}

		if p == "" || p == "." {
			filesIgnored += 1
			continue
		}
		if _, err = os.Stat(p); os.IsNotExist(err) {
			var dir string
			if runtime.GOOS != "windows" {
				dirsplit := strings.Split(p, "/")
				dir = path.Join("/", path.Join(dirsplit[:len(dirsplit)-1]...))
			} else {
				dirwindowsedition := strings.ReplaceAll(p, "/", "\\")
				dirsplit := strings.Split(dirwindowsedition, "\\")
				dir = path.Join(dirsplit[:len(dirsplit)-1]...)
			}
			mkerr := os.MkdirAll(dir, 0700)
			if mkerr != nil {
				return err
			}
		}

		f, err := os.Create(p)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(df.Content)
		if err != nil {
			return err
		}

		fmt.Printf("✅ pulled %v\n", p)
	}

	if filesIgnored > 0 {
		fmt.Printf("🤷 ignored %v files", filesIgnored)
	}

	return nil
}
