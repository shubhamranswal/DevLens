package scanner

import (
	"os"
	"path/filepath"
	"strings"

	"devlens/internal/model"
)

var ignoredDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"dist":         true,
	"build":        true,
	"bin":          true,
}

func Scan(root string) (*model.ScanResult, error) {
	result := &model.ScanResult{
		Extensions:  make(map[string]int),
		Directories: make(map[string]bool),
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		name := info.Name()

		if info.IsDir() {
			if ignoredDirs[name] {
				return filepath.SkipDir
			}
			result.Directories[name] = true
			return nil
		}

		result.Files = append(result.Files, path)

		ext := strings.ToLower(filepath.Ext(name))
		if ext != "" {
			result.Extensions[ext]++
		}

		return nil
	})

	return result, err
}
