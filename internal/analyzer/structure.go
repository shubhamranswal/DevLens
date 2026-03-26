package analyzer

import "devlens/internal/model"

func AnalyzeStructure(scan *model.ScanResult) map[string]string {
	structure := make(map[string]string)

	for dir := range scan.Directories {
		switch dir {
		case "backend", "api":
			structure["/"+dir] = "Backend Logic"

		case "frontend", "ui":
			structure["/"+dir] = "Frontend UI"

		case "scripts":
			structure["/"+dir] = "Automation Scripts"

		case "config":
			structure["/"+dir] = "Configuration"
		}
	}

	return structure
}
