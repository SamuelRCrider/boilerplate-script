package main

// to build: go build -o scriptname

import (
	"fmt"
	"os"
	"strings"

	next_scripts "sam.crider/boilerplate-script/full_stack/next"
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
	next_scripts.Next_Firebase("test-next", "10009")
}

func _main() {

	// parse args
	args := os.Args[1:]

	fmt.Println(args)

	if len(args) > 0 {
		if args[0] == "--help" {
			utils.PrintHelp()
			return
		}
		// for now, we only support the --help flag
		fmt.Println("Currently, we only support the --help flag")
		return
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
