package no_database

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Express_NoAuth_NoDB() {

	// create index.ts file in new project
	utils.Create_File("index.ts", generated.File__index)

	// install cors, dotenv, express, nodemon, ts-node
	cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// install dev deps: cors types, express types, prisma
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "@types/cors", "@types/express")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make app.ts
	utils.Create_File("app.ts", generated.File__noAuthApp)

	utils.Work_wrapper(func() {
		// make utils folder, cd into it
		utils.Mkdir_chdir("utils")

		// create global.d.ts file
		utils.Create_File("global.d.ts", generated.File__expressNoAuthGlobal)

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

		// create service file and types file
		utils.Create_File("controller.ts", generated.File__noAuthController)
		utils.Create_File("types.ts", generated.File__firebaseFrontTypes)

	}, "Creating Utils and Library files...")()
}
