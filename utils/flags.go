package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const helpText = `

Chiks: Your one-stop solution for hatching modern web projects with ease!

Usage: chiks [options]

Options:
  -h, --help           Display this help message

Currently Supported Stacks:
  - Next.js
  - Vite + Express
  - NOTE: Don't see your stack? Please add it!! Check out the CONTRIBUTING.md file! (https://github.com/SamuelRCrider/chiks/CONTRIBUTING.md)

Features:
  - Docker Integration
  - Prisma Database Setup
  - Auth Integration Options
  - Tailwind CSS Integration (optional)

Prerequisites (things you need to have installed):
  - Docker
  - PostgreSQL

For more information, visit: https://github.com/SamuelRCrider/chiks

To update Chiks:
  $ npm update -g chiks

To uninstall Chiks:
  1. $ npm uninstall -g chiks
  2. $ rm $(which chiks)

For bug reports or feature requests, please visit our GitHub repository.
`

func PrintHelp() {
	fmt.Fprint(os.Stderr, helpText)
}

const versionText string = `
Chiks: Your one-stop solution for hatching modern web projects with ease!

Current: user-v

Latest: latest-v

To update Chiks:
  $ npm update -g chiks

`

func PrintVersion() {
	// Read the package.json file
	data, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Println("Error reading package.json:", err)
		os.Exit(1)
	}

	// Parse the JSON data
	var pkg map[string]interface{}
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		fmt.Println("Error parsing package.json:", err)
		os.Exit(1)
	}

	// version user has installed
	currentVersion, ok := pkg["version"].(string)
	if !ok {
		fmt.Println("Error parsing version:", err)
		os.Exit(1)
	}

	// latest version of chiks
	var out bytes.Buffer
	cmd := BoundCommand("npm", "view", "chiks", "version")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	latestVersion := strings.TrimSpace(out.String())

	versionText := Revise_Text(versionText, []Params{{Name: "user-v", Value: currentVersion}, {Name: "latest-v", Value: latestVersion}})
	fmt.Fprint(os.Stderr, versionText)
}
