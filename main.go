package main

// to build: go build -o scriptname

import (
	"fmt"
	"os"
	"strings"

	"sam.crider/boilerplate-script/utils"

	frontend_only_boil "sam.crider/boilerplate-script/frontend_only"
	full_stack_boil "sam.crider/boilerplate-script/full_stack"
)

// stacks is a list of all the stacks that the user can select
var stacks = []string{
	"Vite (Frontend Only)",
	"Next.js (Frontend Only)",
	"Vite + Express",
	"Next.js Full Stack",
	"Add Your Own: (https://github.com/SamuelRCrider/chiks/blob/main/CONTRIBUTING.md)",
}

func main() {

	// parse args
	args := os.Args[1:]

	if len(args) > 0 {
		switch args[0] {
		case "--help":
			utils.PrintHelp()
			return
		case "--version":
			utils.PrintVersion()
			return
		case "-v":
			utils.PrintVersion()
			return
		default:
			fmt.Println("Currently, we only support --help, --version, and -v flags")
			return
		}
	}

	// get the user's selected stack
	stack := utils.Select(
		"Select Your Build Stack:",
		stacks,
	)

	if strings.Contains(stack, "Add Your Own") {
		// link to the contributing guide
		utils.Open("https://github.com/SamuelRCrider/chiks/blob/main/CONTRIBUTING.md")
		return

	}

	// get the user's project name
	project_name := utils.Input(
		"What's the name of this project?",
	)

	// if the stack is frontend only, run the FrontendOnly function
	if strings.Contains(stack, "Frontend") {
		frontend_only_boil.FrontendOnly(stack, project_name)
		return
	}

	docker_check := utils.Select("Docker must be installed and running if you plan on having a database.", []string{
		"Docker is installed and running.",
		"I need to install Docker.",
		"Open Docker for me.",
		"I don't plan on using a database.",
	})
	if strings.Contains(docker_check, "need") {
		utils.Open("https://docs.docker.com/get-docker/")
		return
	} else if strings.Contains(docker_check, "Open") {
		cmd := utils.BoundCommand("open", "-a", "Docker")
		if err := cmd.Run(); err != nil {
			fmt.Println("Docker failed to open, please open it manually and try again.")
			return
		}
	}

	// get users auth preference
	auth_integration := utils.Select(
		"Pick an auth integration:",
		[]string{
			"Firebase",
			"Clerk",
			"None",
		},
	)

	// get the docker port
	docker_port := utils.GetDockerPort()

	// full stack
	full_stack_boil.Full_Stack(stack, project_name, docker_port, auth_integration)
}
