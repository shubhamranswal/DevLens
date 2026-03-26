package model

type ScanResult struct {
	Files       []string
	Extensions  map[string]int
	Directories map[string]bool
}

type AnalysisResult struct {
	ProjectType  string
	TechStack    map[string]string
	Structure    map[string]string
	EntryPoints  []string
	Observations []string
	Frameworks   []string
}
