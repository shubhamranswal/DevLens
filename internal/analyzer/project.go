package analyzer

import (
	"strings"

	"devlens/internal/model"
)

func DetectProjectType(scan *model.ScanResult) string {

	hasGo := scan.Extensions[".go"] > 0
	hasPython := scan.Extensions[".py"] > 0
	hasJS := scan.Extensions[".js"] > 0 || scan.Extensions[".ts"] > 0
	hasReact := scan.Extensions[".jsx"] > 0 || scan.Extensions[".tsx"] > 0
	hasJava := scan.Extensions[".java"] > 0
	hasKotlin := scan.Extensions[".kt"] > 0
	hasDart := scan.Extensions[".dart"] > 0
	hasRust := scan.Extensions[".rs"] > 0

	for _, file := range scan.Files {
		if strings.HasSuffix(file, "main.go") ||
			strings.HasSuffix(file, "main.rs") ||
			strings.HasSuffix(file, "main.py") {
			return "CLI Tool"
		}
	}

	if hasDart && containsDir(scan, "lib") {
		return "Mobile Application (Flutter)"
	}

	if hasJava || hasKotlin {
		if containsDir(scan, "android") {
			return "Mobile Application (Android)"
		}
		return "Backend Service (Java/Kotlin)"
	}

	if hasReact {
		return "Frontend Application (React)"
	}

	if hasJS {
		return "Web Application"
	}

	if hasGo {
		return "Backend Service (Go)"
	}

	if hasPython {
		return "Backend Service (Python)"
	}

	if hasRust {
		return "Backend Service (Rust)"
	}

	return "Unknown"
}

func containsDir(scan *model.ScanResult, name string) bool {
	for dir := range scan.Directories {
		if dir == name {
			return true
		}
	}
	return false
}
