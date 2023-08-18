package config

import (
	"errors"

	"github.com/Tesohh/dotdepot/cli/db"
)

var ErrNoConfigFile = errors.New("no config file found: please create a config.yml in the dotdepot folder (~/.config/dotdepot)")

type Config struct {
	Files []db.Paths `yaml:"files"`
}
