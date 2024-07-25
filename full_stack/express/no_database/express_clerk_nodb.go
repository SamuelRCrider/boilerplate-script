package no_database

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Express_Clerk_NoDB() {

	// create index.ts file in new project
	utils.Create_File("index.ts", generated.File__index)

	// install cors, dotenv, express, nodemon, ts-node
	cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node", "@clerk/clerk-sdk-node")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// install dev deps: cors types, express types
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "@types/cors", "@types/express", "@clerk/types")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make app.ts
	utils.Create_File("app.ts", generated.File__expressClerkApp)

	// make env file
	utils.Create_File(".env", generated.File__expressClerkEnvNoDb)

	utils.Work_wrapper(func() {
		// make utils folder, cd into it
		utils.Mkdir_chdir("utils")

		// create global.d.ts file
		utils.Create_File("global.d.ts", generated.File__expressClerkGlobal)

		// cd out of utils
		err := os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir lib
		utils.Mkdir_chdir("lib")

		// mkdir controllers
		utils.Mkdir_chdir("controllers")

		// mkdir users
		utils.Mkdir_chdir("users")

		// create controller and types files
		utils.Create_File("controller.ts", generated.File__expressClerkController)
		utils.Create_File("types.ts", generated.File__firebaseAuthTypes)

	}, "Creating Utils and Library files...")()
}
