package analyzer

import (
	"strings"

	"devlens/internal/model"
)

func DetectEntryPoints(scan *model.ScanResult) []string {
	var entries []string

	for _, file := range scan.Files {
		switch {
		case strings.HasSuffix(file, "main.go"),
			strings.HasSuffix(file, "app.py"),
			strings.HasSuffix(file, "index.js"),
			strings.HasSuffix(file, "App.jsx"):

			entries = append(entries, file)
		}
	}

	return entries
}
