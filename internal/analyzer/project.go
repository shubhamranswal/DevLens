package analyzer

import "devlens/internal/model"

func DetectProjectType(scan *model.ScanResult) string {

	if scan.Directories["cmd"] && scan.Extensions[".go"] > 0 {
		return "CLI Tool"
	}

	if scan.Extensions[".jsx"] > 0 || scan.Extensions[".tsx"] > 0 {
		return "Frontend Application"
	}

	if scan.Extensions[".py"] > 0 || scan.Extensions[".go"] > 0 {
		return "Backend Service"
	}

	return "Unknown"
}
