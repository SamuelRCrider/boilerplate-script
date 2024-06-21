package main

// to build: go build -o scriptname

import (
	"fmt"
	"os"

	express_boil "sam.crider/boilerplate-script/express"
	"sam.crider/boilerplate-script/utils"
	vite_boil "sam.crider/boilerplate-script/vite"
)

func main() {
	// get the user's selected stack
	stack := utils.Select_Stack(
		"Select Your Build Stack:",
		[]string{
			"Vite + Express",
			"Next.js",
		},
	)

	// get the user's project name
	project_name := utils.Name_Project(
		"What's the name of this project?",
	)

	if stack == "Vite + Express" {
		// create a directory for the project, 0755 is the permission bits
		err := os.Mkdir(project_name, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into the new project
		err = os.Chdir(project_name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// initialize git for the project
		cmd := utils.BoundCommand("git", "init")
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// create the frontend
		vite_boil.Vite()

		// create the backend
		express_boil.Express()

		fmt.Println("Success! Boilerplate created.")
		return
	}
	if stack == "Next.js" {
		fmt.Println("Coming Soon")
		fmt.Println("Success! Boilerplate created.")
		return
	}

	fmt.Println("Failure. Maybe you didn't select an option?")
}
