package main

import (
	"devlens/internal/analyzer"
	"devlens/internal/model"
	"devlens/internal/output"
	"devlens/internal/scanner"
	"fmt"
	"os"
	"strings"
)

const version = "v0.1.0"

func main() {

	args := os.Args[1:]

	var jsonOutput bool
	var versionFlag bool
	var helpFlag bool
	var fastMode bool

	var filteredArgs []string

	for _, arg := range args {
		switch arg {
		case "--json":
			jsonOutput = true
		case "--version":
			versionFlag = true
		case "--help":
			helpFlag = true
		case "--fast":
			fastMode = true
		default:
			filteredArgs = append(filteredArgs, arg)
		}
	}

	if versionFlag {
		fmt.Println("DevLens", version)
		return
	}

	if helpFlag {
		fmt.Println(`
DevLens - Codebase Analyzer

Usage:
  devlens analyze <path>

Flags:
  --json       Output JSON
  --version    Show version
  --help       Show help
`)
		return
	}

	if len(filteredArgs) < 2 || filteredArgs[0] != "analyze" {
		fmt.Println("Usage: devlens analyze <path> [--json]")
		return
	}

	path := filteredArgs[1]

	var scan *model.ScanResult

	if strings.HasPrefix(path, "http") && fastMode {
		scan = analyzer.RemoteScan(path)
		if scan == nil {
			fmt.Println("Failed to analyze remote repository")
			return
		}
	} else {
		var err error
		scan, err = scanner.Scan(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	result := &model.AnalysisResult{
		ProjectType:  analyzer.DetectProjectType(scan),
		TechStack:    analyzer.DetectTech(scan),
		Structure:    analyzer.AnalyzeStructure(scan),
		EntryPoints:  analyzer.DetectEntryPoints(scan),
		Observations: analyzer.GenerateObservations(scan),
		Frameworks:   analyzer.DetectFrameworks(scan),
	}

	if result.Observations == nil {
		result.Observations = []string{}
	}
	if result.Frameworks == nil {
		result.Frameworks = []string{}
	}

	output.Print(result, jsonOutput)
}
