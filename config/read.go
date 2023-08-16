package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func Read[T any](filename string) (*T, error) {
	cfgdir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dir := path.Join(cfgdir, "/.config/dotdepot", filename)
	f, err := os.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	var file T
	yaml.Unmarshal(f, &file)
	return &file, nil
}
