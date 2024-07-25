package frontend_only_boil

import (
	"fmt"

	frontend_only_scripts "sam.crider/boilerplate-script/frontend_only/frontend_scripts"

	"sam.crider/boilerplate-script/utils"
)

func FrontendOnly(stack string, project_name string) {
	// auth integration
	auth_integration := utils.Select(
		"Pick an auth integration:",
		[]string{
			"Clerk",
			"None",
		},
	)

	// switch case for stack
	switch stack {
	case "Vite (Frontend Only)":

		// switch case for auth integration
		switch auth_integration {
		case "Clerk":
			// create the frontend
			frontend_only_scripts.Vite_FrontendClerk(project_name)
		case "None":
			// create the frontend
			frontend_only_scripts.Vite_FrontendNoAuth(project_name)
		default:
			fmt.Println("Failure. Maybe you didn't select an option?")
			return
		}

		fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
		return

	case "Next.js (Frontend Only)":
		// switch case for auth integration
		switch auth_integration {
		case "Clerk":
			// create the frontend
			frontend_only_scripts.Next_FrontendClerk(project_name)
		case "None":
			// create the frontend
			frontend_only_scripts.Next_FrontendNoAuth(project_name)
		default:
			fmt.Println("Failure. Maybe you didn't select an option?")
			return
		}

		fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
		return

	default:
		fmt.Println("Failure. Maybe you didn't select an option?")
		return
	}

}
