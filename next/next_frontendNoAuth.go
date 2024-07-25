package next_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/utils"
)

func Next_FrontendNoAuth(project_name string) {

	// create next app
	cmd := utils.BoundCommand("npx", "create-next-app", project_name, "--typescript")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// cd into project
	err := os.Chdir(project_name)
	if err != nil {
		fmt.Println(err)
		return
	}

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

	// remove readme and replace with no auth readme
	err = os.Remove("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("README.md", generated.File__nextNoAuthFrontReadMe)

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

		if ui_check == "Shadcn UI" {
			// mkdir lib
			utils.Mkdir_chdir("lib")

			// make utils file
			utils.Create_File("utils.ts", generated.File__viteShadcnUtils)
		}

	}, "Creating Utils, Components, and Library folders...")()
}
