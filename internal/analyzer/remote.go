package analyzer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"devlens/internal/model"
)

type GitHubFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func RemoteScan(repoURL string) *model.ScanResult {

	owner, repo := parseGitHubURL(repoURL)
	if owner == "" || repo == "" {
		fmt.Println("Invalid GitHub URL")
		return nil
	}

	result := &model.ScanResult{
		Files:       []string{},
		Extensions:  make(map[string]int),
		Directories: make(map[string]bool),
	}

	err := fetchDir(owner, repo, "", result)
	if err != nil {
		fmt.Println("Failed to fetch repo:", err)
		return nil
	}

	return result
}

func fetchDir(owner, repo, path string, result *model.ScanResult) error {

	apiURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s",
		owner, repo, path,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var files []GitHubFile
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return err
	}

	for _, f := range files {
		if f.Type == "dir" {
			dirParts := strings.Split(f.Path, "/")
			for _, d := range dirParts {
				result.Directories[d] = true
			}
			err := fetchDir(owner, repo, f.Path, result)
			if err != nil {
				continue
			}
		} else {
			result.Files = append(result.Files, f.Path)
			ext := getExtension(f.Name)
			if ext != "" {
				result.Extensions[ext]++
			}
		}
	}

	return nil
}

func parseGitHubURL(url string) (string, string) {
	parts := strings.Split(url, "/")
	if len(parts) < 5 {
		return "", ""
	}
	return parts[3], parts[4]
}

func getExtension(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) > 1 {
		return "." + parts[len(parts)-1]
	}
	return ""
}
