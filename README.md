# DevLens

DevLens is a CLI tool that analyzes a codebase and provides a structured summary of its architecture, tech stack, and organization.

## Usage

```bash
devlens analyze .
````

## What It Does

* Scans project structure
* Detects technologies via heuristics
* Infers project type
* Identifies key directories
* Finds entry points
* Generates insights

## Design Philosophy

* Simple heuristics > complex parsing
* Fast understanding > perfect accuracy
* Zero dependencies (std lib only)

## Limitations

* No deep code parsing
* Heuristic-based (may miss edge cases)
* No framework-level intelligence (yet)

## Future Ideas

* Framework detection (React, FastAPI, etc.)
* Config parsing
* Dependency graph

