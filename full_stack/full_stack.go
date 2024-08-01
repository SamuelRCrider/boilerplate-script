package full_stack_boil

import (
	"fmt"
	"os"

	"sam.crider/boilerplate-script/utils"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	express_scripts "sam.crider/boilerplate-script/full_stack/express"
	next_scripts "sam.crider/boilerplate-script/full_stack/next"
	vite_scripts "sam.crider/boilerplate-script/full_stack/vite"
)

func Full_Stack(stack string, project_name string, docker_port string, auth_integration string) {
	// switch case for stack
	switch stack {
	case "Vite + Express":
		utils.Work_wrapper(func() {
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
		}, "Initializing project...")()

		// TODO: make this a switch case
		if auth_integration == "Firebase" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressFirebaseReadme)

			// create the frontend
			vite_scripts.Vite_FirebaseAuth()

			// create the backend
			express_scripts.Express_FirebaseAuth(docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "Clerk" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressClerkReadme)

			// create the frontend
			vite_scripts.Vite_ClerkAuth()

			// create the backend
			express_scripts.Express_ClerkAuth(project_name, docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "None" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressNoAuthReadme)

			// create the frontend
			vite_scripts.Vite_NoAuth()

			// create the backend
			express_scripts.Express_NoAuth(docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return
	case "Next.js Full Stack":
		// TODO: make this a switch case
		if auth_integration == "Firebase" {
			// create the app
			next_scripts.Next_Firebase(project_name, docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "Clerk" {
			// create the app
			next_scripts.Next_ClerkAuth(project_name, docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "None" {
			// create the app
			next_scripts.Next_NoAuth(project_name, docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return
	default:
		fmt.Println("Failure. Maybe you didn't select an option? Or maybe you clicked Add Your Own. In that case, check out: https://github.com/SamuelRCrider/chiks/blob/main/CONTRIBUTING.md")
		return
	}

}
