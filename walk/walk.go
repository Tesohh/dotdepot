package walk

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func Walk(basedir string, justFileName bool) ([]string, error) {
	paths := make([]string, 0)
	err := filepath.WalkDir(basedir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if justFileName {
				paths = append(paths, strings.Replace(path, basedir, "", 1))
			} else {
				paths = append(paths, path)
			}
		}
		return nil
	})
	return paths, err
}
