package no_database

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Next_Firebase_NoDB() {

	// check if there is a tailwind.config.js file
	if _, err := os.Stat("tailwind.config.ts"); err == nil {
		// if there is, ask if user wants to use daisyUI, shadcn UI, or just Tailwind
		ui_check = utils.Select(
			"Which UI framework would you like to use?",
			[]string{
				"Shadcn UI",
				"DaisyUI",
				"None (base Tailwind)",
			},
		)

		// TODO: make this a switch case if we want to add more UI frameworks
		if ui_check == "Shadcn UI" {
			// install tailwind with shadcn ui
			// Run the shadcn-ui init command
			cmd := utils.BoundCommand("npx", "shadcn-ui@latest", "init")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// remove the components.json file
			err = os.Remove("components.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			// replace the components.json file
			utils.Create_File("components.json", generated.File__nextComponentsJson)

			// remove the lib folder
			err = os.RemoveAll("src/lib")
			if err != nil {
				fmt.Println(err)
				return
			}

			// remove the components folder
			err = os.RemoveAll("src/components")
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if ui_check == "DaisyUI" {
			// install tailwind with daisy ui
			// Run the daisy init command
			cmd := utils.BoundCommand("npm", "i", "-D", "daisyui")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// replace the tailwind.config.ts file
			err = os.Remove("tailwind.config.ts")
			if err != nil {
				fmt.Println(err)
				return
			}
			utils.Create_File("tailwind.config.ts", generated.File__nextTailwindConfig)

		}
	}

	// remove readme and replace with firebase readme
	err := os.Remove("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("README.md", generated.File__nextFirebaseReadme)

	// install deps
	cmd_deps := utils.BoundCommand("npm", "install", "firebase")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// replace the .env file
	err = os.Remove(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.Create_File(".env", generated.File__nextFirebaseEnvNoDb)

	// replace the gitignore file
	err = os.Remove(".gitignore")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File(".gitignore", generated.File__nextGitignore)

	utils.Work_wrapper(func() {
		// cd src directory
		err = os.Chdir("src")
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into app
		err = os.Chdir("app")
		if err != nil {
			fmt.Println(err)
			return
		}

		// remove page.tsx file
		err = os.Remove("page.tsx")
		if err != nil {
			fmt.Println(err)
			return
		}

		// replace page.tsx file
		utils.Create_File("page.tsx", generated.File__nextFirebasePage)

		// remove the layout file
		err = os.Remove("layout.tsx")
		if err != nil {
			fmt.Println(err)
			return
		}

		// replace the layout file
		utils.Create_File("layout.tsx", generated.File__nextFirebaseLayout)

		// mkdir login
		utils.Mkdir_chdir("login")

		// make page.tsx file
		utils.Create_File("page.tsx", generated.File__nextFirebaseLoginPage)

		// cd out of login
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir dashboard
		utils.Mkdir_chdir("dashboard")

		// make page.tsx file
		utils.Create_File("page.tsx", generated.File__nextFirebaseDashboardPage)

		// cd out of app
		err = os.Chdir("../../")

		// mkdir hooks
		utils.Mkdir_chdir("hooks")

		// make useAuth hook
		utils.Create_File("useAuth.ts", generated.File__nextFirebaseHook)

		// cd out of hooks
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir components
		utils.Mkdir_chdir("components")

		if ui_check == "Shadcn UI" {
			// mkdir shadcn
			err = os.Mkdir("shadcn", 0755)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// mkdir pages
		utils.Mkdir_chdir("pages")

		// create example component
		utils.Create_File("Example.tsx", generated.File__exampleComponent)

		// cd out of pages
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir compound
		utils.Mkdir_chdir("compound")

		// create example component
		utils.Create_File("Example.tsx", generated.File__exampleComponent)

		// cd out of compound
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir base
		utils.Mkdir_chdir("base")

		// create example component
		utils.Create_File("Example.tsx", generated.File__exampleComponent)

		// cd out of components
		err = os.Chdir("../../")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir providers
		utils.Mkdir_chdir("providers")

		// make AuthProvider file
		utils.Create_File("AuthProvider.tsx", generated.File__nextFirebaseProvider)

		// cd out of providers
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir lib
		utils.Mkdir_chdir("lib")

		if ui_check == "Shadcn UI" {
			// make utils file
			utils.Create_File("utils.ts", generated.File__viteShadcnUtils)
		}

		// mkdir middleware
		utils.Mkdir_chdir("middleware")

		// make middleware file
		utils.Create_File("middleware.ts", generated.File__nextFirebaseMiddleware)

		// cd out of middleware
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir firebase
		utils.Mkdir_chdir("firebase")

		// create firebase config file
		utils.Create_File("config.ts", generated.File__nextFirebaseConfig)

	}, "Creating Utils, Components, and Library folders...")()
}
