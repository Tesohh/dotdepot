package config

import "github.com/Tesohh/dotdepot/cli/db"

type Config struct {
	Files []db.Paths `yaml:"files"`
}
