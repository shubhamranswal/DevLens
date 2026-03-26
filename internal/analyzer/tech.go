package analyzer

import (
	"strings"

	"devlens/internal/model"
)

func DetectTech(scan *model.ScanResult) map[string]string {
	stack := make(map[string]string)

	for _, file := range scan.Files {
		switch {
		case strings.HasSuffix(file, "go.mod"):
			stack["Backend"] = "Go"

		case strings.HasSuffix(file, "package.json"):
			stack["Frontend"] = "Node.js"

		case strings.HasSuffix(file, "requirements.txt"):
			stack["Backend"] = "Python"

		case strings.HasSuffix(file, "Dockerfile"):
			stack["DevOps"] = "Docker"

		case strings.HasSuffix(file, "firebase.json"):
			stack["Backend"] = "Firebase"
		}
	}

	if _, ok := scan.Extensions[".env"]; ok {
		stack["Config"] = "Environment Variables"
	}

	return stack
}
