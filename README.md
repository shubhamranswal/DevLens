# DevLens

A CLI tool to quickly understand unfamiliar codebases through structured analysis.

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Version](https://img.shields.io/badge/version-v0.1.0-black.svg)]()
[![CLI Tool](https://img.shields.io/badge/type-CLI-lightgrey.svg)]()


DevLens is a lightweight CLI tool that analyzes a codebase and provides a structured summary of its technology stack, architecture, and organization.

It is designed to help developers quickly understand unfamiliar repositories without manual exploration.

---

## Installation

### Option 1: Install globally

```bash
go install ./cmd/devlens
````

Ensure your Go bin directory is added to PATH:

```
C:\Users\<your-username>\go\bin
```

---

### Option 2: Build locally

```bash
go build -o devlens.exe ./cmd/devlens
```

---

## Usage

### Analyze a project

```bash
devlens analyze .
```

### Analyze a remote repository (fast mode)

```bash
devlens analyze https://github.com/user/repo --fast
````

* Uses GitHub API (no cloning required)
* Faster for large repositories
* May have limited accuracy for framework detection

### JSON output

```bash
devlens analyze . --json
```

or 

```bash
devlens analyze . --json > report.json
```

### Other commands

```bash
devlens --version
devlens --help
```

---

## Features

### Technology Detection

Identifies technologies using file patterns such as:

* `go.mod` → Go
* `package.json` → Node.js
* `requirements.txt` → Python
* `Dockerfile` → Docker
* `firebase.json` → Firebase

---

### Framework Detection

Detects commonly used frameworks, including:

* React, Next.js, Vite
* FastAPI, Django, Flask
* Express.js

---

### Project Type Inference

Infers project type based on structure and files:

* CLI Tool
* Backend Service
* Frontend Application

---

### Structure Analysis

Identifies key directories and their roles:

* `/backend`, `/api` → backend logic
* `/frontend`, `/ui` → frontend layer
* `/scripts` → automation
* `/config` → configuration

---

### Entry Point Detection

Detects likely entry points such as:

* `main.go`
* `app.py`
* `index.js`
* `App.jsx`

---

### Observations

Provides heuristic insights such as:

* Dockerized setup detected
* Environment variables used
* Monorepo-like structure
* Presence of test files

---

## Example Output

### Standard Output

```
DevLens Analysis

Project Type:
  CLI Tool

Tech Stack:
 - Backend: Go

Frameworks:

Entry Points:
 - cmd/devlens/main.go
```

---

### JSON Output

```json
{
  "project_type": "CLI Tool",
  "tech_stack": {
    "Backend": "Go"
  },
  "frameworks": [],
  "entry_points": [
    "cmd/devlens/main.go"
  ]
}
```

---

## Design Philosophy

* Prefer simple heuristics over complex parsing
* Optimize for speed and clarity
* Avoid external dependencies
* Keep behavior predictable and transparent

---

## Limitations

* Heuristic-based detection may not cover all edge cases
* No deep code parsing
* Framework detection is best-effort

---

## Roadmap

* Parallel scanning for performance improvement
* Enhanced architecture detection
* Project complexity analysis
* Plugin system
* Editor integrations

---

## Contributing

Contributions are welcome. Please open an issue or submit a pull request for improvements or feature suggestions.

---

## License

MIT

