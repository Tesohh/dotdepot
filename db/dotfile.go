package db

import (
	"errors"
	"os"
	"path"
	"runtime"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrCouldntDetermineOS = errors.New("couldn't determine your current os. what kind of computer are you running man")

type Paths struct {
	Windows string `bson:"windows,omitempty"`
	MacOS   string `bson:"macos,omitempty"`
	Linux   string `bson:"linux,omitempty"`
}

func (p Paths) ToQuery() map[string]string {
	m := make(map[string]string)
	if p.Windows != "" {
		m["windows"] = p.Windows
	}
	if p.MacOS != "" {
		m["macos"] = p.MacOS
	}
	if p.Linux != "" {
		m["linux"] = p.Linux
	}
	return m
}

func (p Paths) ForCurrentOS() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	var fpath string
	switch runtime.GOOS {
	case "windows":
		fpath = p.Windows
	case "darwin":
		fpath = p.MacOS
	case "linux":
		fpath = p.Linux
	default:
		return "", ErrCouldntDetermineOS
	}
	fpath = strings.Replace(fpath, "~", home, 1)
	fpath = path.Clean(fpath)
	return fpath, nil
}

type Dotfile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FileName    string             `bson:"filename,omitempty"`
	UserName    string             `bson:"username,omitempty"`
	Content     string             `bson:"content,omitempty"`
	Paths       Paths              `bson:"paths,omitempty"`
	IsDirectory bool               `bson:"isdir,omitempty"`
}

func (df Dotfile) IsEmpty() bool {
	return df == Dotfile{}
}
