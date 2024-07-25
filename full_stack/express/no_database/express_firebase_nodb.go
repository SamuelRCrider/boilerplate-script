package no_database

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Express_Firebase_NoDB() {

	// create index.ts file in new project
	utils.Create_File("index.ts", generated.File__index)

	// install cors, dotenv, express, nodemon, ts-node
	cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node", "firebase-admin")

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
	utils.Create_File("app.ts", generated.File__firebaseAuthApp)

	// make firebase service account key file
	utils.Create_File("serviceAccountKey.json", generated.File__serviceAccountKey)

	// replace the gitignore file
	err := os.Remove(".gitignore")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File(".gitignore", generated.File__firebaseGitignore)

	utils.Work_wrapper(func() {
		// make utils folder, cd into it
		utils.Mkdir_chdir("utils")

		// create requireAuth.ts
		utils.Create_File("requireAuth.ts", generated.File__firebaseRequireAuth)

		// create global.d.ts
		utils.Create_File("global.d.ts", generated.File__expressFirebaseGlobal)

		// cd out of utils
		err := os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir lib
		utils.Mkdir_chdir("lib")

		// mkdir firebase
		utils.Mkdir_chdir("firebase")

		// create firebase config file
		utils.Create_File("config.ts", generated.File__firebaseConfig)

		// cd out of firebase
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir controllers
		utils.Mkdir_chdir("controllers")

		// mkdir auth
		utils.Mkdir_chdir("auth")

		// create controller and types files
		utils.Create_File("controller.ts", generated.File__firebaseAuthController)
		utils.Create_File("types.ts", generated.File__firebaseFrontTypes)

	}, "Creating Library files...")()
}
