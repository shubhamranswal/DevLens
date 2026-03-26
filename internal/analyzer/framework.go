package analyzer

import (
	"os"
	"strings"

	"devlens/internal/model"
)

func DetectFrameworks(scan *model.ScanResult) []string {
	var frameworks []string

	for _, file := range scan.Files {

		switch {

		// Node ecosystem
		case strings.HasSuffix(file, "package.json"):
			data, err := os.ReadFile(file)
			if err == nil {
				content := string(data)

				if strings.Contains(content, "react") {
					frameworks = append(frameworks, "React")
				}
				if strings.Contains(content, "next") {
					frameworks = append(frameworks, "Next.js")
				}
				if strings.Contains(content, "vite") {
					frameworks = append(frameworks, "Vite")
				}
				if strings.Contains(content, "express") {
					frameworks = append(frameworks, "Express.js")
				}
			}

		// Python ecosystem
		case strings.HasSuffix(file, "requirements.txt"):
			data, err := os.ReadFile(file)
			if err == nil {
				content := string(data)

				if strings.Contains(content, "fastapi") {
					frameworks = append(frameworks, "FastAPI")
				}
				if strings.Contains(content, "django") {
					frameworks = append(frameworks, "Django")
				}
				if strings.Contains(content, "flask") {
					frameworks = append(frameworks, "Flask")
				}
			}
		}
	}

	return unique(frameworks)
}

func unique(input []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
