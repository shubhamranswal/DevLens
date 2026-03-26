package output

import (
	"encoding/json"
	"fmt"

	"devlens/internal/model"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Green  = "\033[32m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
)

func Print(result *model.AnalysisResult, asJSON bool) {

	if asJSON {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
		return
	}

	fmt.Println(Bold + Cyan + "DevLens Analysis" + Reset)

	fmt.Println("\n" + Bold + "Project Type:" + Reset)
	fmt.Println(" ", Green+result.ProjectType+Reset)

	fmt.Println("\n" + Bold + "Tech Stack:" + Reset)
	for k, v := range result.TechStack {
		fmt.Printf(" - %s: %s%s%s\n", k, Yellow, v, Reset)
	}

	fmt.Println("\nFrameworks:")
	for _, f := range result.Frameworks {
		fmt.Println(" -", f)
	}

	fmt.Println("\n" + Bold + "Structure:" + Reset)
	for k, v := range result.Structure {
		fmt.Printf(" - %s → %s\n", k, v)
	}

	fmt.Println("\n" + Bold + "Entry Points:" + Reset)
	for _, e := range result.EntryPoints {
		fmt.Println(" -", e)
	}

	fmt.Println("\n" + Bold + "Observations:" + Reset)
	for _, o := range result.Observations {
		fmt.Println(" -", o)
	}
}
