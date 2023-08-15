package auth

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Read() (Credentials, error) {
	cfgdir, err := os.UserHomeDir()
	if err != nil {
		return Credentials{}, err
	}
	dir := path.Join(cfgdir, "/.config/dotdepot", "login.yml")
	f, err := os.ReadFile(dir)
	if err != nil {
		return Credentials{}, err
	}
	creds := Credentials{}
	yaml.Unmarshal(f, &creds)
	return creds, nil
}
