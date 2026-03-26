package main

import (
	"fmt"
	"os"

	"devlens/internal/analyzer"
	"devlens/internal/model"
	"devlens/internal/output"
	"devlens/internal/scanner"
)

func main() {

	if len(os.Args) < 3 || os.Args[1] != "analyze" {
		fmt.Println("Usage: devlens analyze <path>")
		return
	}

	path := os.Args[2]

	scan, err := scanner.Scan(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := &model.AnalysisResult{
		ProjectType:  analyzer.DetectProjectType(scan),
		TechStack:    analyzer.DetectTech(scan),
		Structure:    analyzer.AnalyzeStructure(scan),
		EntryPoints:  analyzer.DetectEntryPoints(scan),
		Observations: analyzer.GenerateObservations(scan),
	}

	output.Print(result)
}
