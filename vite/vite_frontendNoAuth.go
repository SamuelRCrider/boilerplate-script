package vite_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/utils"
)

func Vite_FrontendNoAuth(project_name string) {

	cmd := utils.BoundCommand("npx", "create-vite@latest", project_name)

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// cd into frontend
	err := os.Chdir(project_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// npm install all the vite packages
	cmd = utils.BoundCommand("npm", "install")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// create .env file
	utils.Create_File(".env", generated.File__noAuthFrontEnv)

	// replace the gitignore file
	err = os.Remove(".gitignore")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File(".gitignore", generated.File__firebaseFrontGitignore)

	// replace the readme file
	err = os.Remove("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("README.md", generated.File__viteFrontendNoAuthReadme)

	// ask if user wants tailwind
	tailwind_check := utils.Select(
		"Do you need tailwind for this project?",
		[]string{
			"Yes",
			"No",
		},
	)

	if tailwind_check == "Yes" {

		// ask if user would like to add daisyUI, Shadcn UI, or just Tailwind
		ui_check := utils.Select(
			"Which UI framework would you like to use?",
			[]string{
				"Shadcn UI",
				"DaisyUI",
				"None (base Tailwind)",
			},
		)

		/* install tailwind */

		cmd := utils.BoundCommand("npm", "install", "-D", "tailwindcss", "postcss", "autoprefixer")

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// initialize tailwind
		cmd = utils.BoundCommand("npx", "tailwindcss", "init", "-p")

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// replace tailwind config file
		err := os.Remove("tailwind.config.js")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File("tailwind.config.js", generated.File__firebaseFrontTailwindConfig)

		// replace index.css file
		// cd into src
		err = os.Chdir("src")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = os.Remove("index.css")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File("index.css", generated.File__firebaseFrontIndexCss)

		// cd out of src
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// TODO: make this a switch case if we want to add more UI frameworks
		if ui_check == "Shadcn UI" {
			// install tailwind with shadcn ui

			// update the tsconfig.json file
			err = os.Remove("tsconfig.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			utils.Create_File("tsconfig.json", generated.File__viteTsconfig)

			// install node types
			cmd = utils.BoundCommand("npm", "i", "-D", "@types/node")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// update the vite.config.ts file
			err = os.Remove("vite.config.ts")
			if err != nil {
				fmt.Println(err)
				return
			}
			utils.Create_File("vite.config.ts", generated.File__viteShadConfig)

			// update tsconfig.app.json file
			err = os.Remove("tsconfig.app.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			utils.Create_File("tsconfig.app.json", generated.File__viteTsconfigApp)

			// inform user of shadcn ui init
			choice = utils.Select(
				"Shadcn UI is about to ask you a bunch of questions. Choose all the defaults and let Chiks configure it for you!",
				[]string{
					"Ok! - Currently, this is your only choice",
				},
			)

			// run shadcn ui init
			cmd = utils.BoundCommand("npx", "shadcn-ui@latest", "init")
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
			utils.Create_File("components.json", generated.File__viteComponentsJson)

			// remove the components folder
			err = os.RemoveAll("src/components")
			if err != nil {
				fmt.Println(err)
				return
			}

			// remove the app folder
			err = os.RemoveAll("app")
			if err != nil {
				fmt.Println(err)
				return
			}

			// update index.css file
			err = os.Remove("src/index.css")
			if err != nil {
				fmt.Println(err)
				return
			}
			utils.Create_File("src/index.css", generated.File__viteShadcnIndex)

			// remove the lib folder
			err = os.RemoveAll("src/lib")
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if ui_check == "DaisyUI" {
			// install tailwind with daisy ui

			// install daisy ui
			cmd = utils.BoundCommand("npm", "i", "-D", "daisyui")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// remove the tailwind.config.js file
			err = os.Remove("tailwind.config.js")
			if err != nil {
				fmt.Println(err)
				return
			}

			// replace the tailwind.config.js file
			utils.Create_File("tailwind.config.js", generated.File__viteDaisyTconfig)

		}

	}

	// cd into src
	err = os.Chdir("src")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir components
	utils.Mkdir_chdir("components")

	if choice == "Ok! - Currently, this is your only choice" {
		// mkdir shadcn
		err = os.Mkdir("shadcn", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// mkdir pages
	utils.Mkdir_chdir("pages")

	// make example component
	utils.Create_File("Example.tsx", generated.File__exampleComponent)

	// cd out of pages
	err = os.Chdir("..")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir compound
	utils.Mkdir_chdir("compound")

	// make example component
	utils.Create_File("Example.tsx", generated.File__exampleComponent)

	// cd out of compound
	err = os.Chdir("..")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir base
	utils.Mkdir_chdir("base")

	// make example component
	utils.Create_File("Example.tsx", generated.File__exampleComponent)

	// cd out of components
	err = os.Chdir("../../")
	if err != nil {
		fmt.Println(err)
		return
	}

	if choice == "Ok! - Currently, this is your only choice" {
		// mkdir lib
		utils.Mkdir_chdir("lib")

		// make utils file
		utils.Create_File("utils.ts", generated.File__viteShadcnUtils)
	}

}
