package output

import (
	"fmt"

	"devlens/internal/model"
)

func Print(result *model.AnalysisResult) {
	fmt.Println("🔍 DevLens Analysis\n")

	fmt.Println("Project Type:")
	fmt.Println(" ", result.ProjectType)

	fmt.Println("\nTech Stack:")
	for k, v := range result.TechStack {
		fmt.Printf(" - %s: %s\n", k, v)
	}

	fmt.Println("\nStructure:")
	for k, v := range result.Structure {
		fmt.Printf(" - %s → %s\n", k, v)
	}

	fmt.Println("\nEntry Points:")
	for _, e := range result.EntryPoints {
		fmt.Println(" -", e)
	}

	fmt.Println("\nObservations:")
	for _, o := range result.Observations {
		fmt.Println(" -", o)
	}
}
