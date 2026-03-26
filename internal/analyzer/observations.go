package analyzer

import "devlens/internal/model"

func GenerateObservations(scan *model.ScanResult) []string {
	var obs []string

	if scan.Directories["docker"] {
		obs = append(obs, "Dockerized setup detected")
	}

	if scan.Extensions[".env"] > 0 {
		obs = append(obs, "Environment variables used")
	}

	if scan.Directories["packages"] || scan.Directories["services"] {
		obs = append(obs, "Monorepo-like structure")
	}

	if scan.Extensions[".test"] > 0 || scan.Extensions["_test.go"] > 0 {
		obs = append(obs, "Contains test files")
	}

	return obs
}
