package express_scripts

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/full_stack/express/no_database"
	"sam.crider/boilerplate-script/utils"
)

func Express_FirebaseAuth(docker_port string) {
	// mkdir for backend, 0755 is the permission bits
	utils.Mkdir_chdir("backend")

	// initialize npm project
	cmd_npm := utils.BoundCommand("npm", "init")
	if err := cmd_npm.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// ask if user needs a database
	db_check := utils.Select(
		"Do you need a database?",
		[]string{
			"Yes - PostgreSQL",
			"No",
		},
	)

	if db_check == "No" {
		// build backend without a database
		no_database.Express_Firebase_NoDB()
		return
	}

	// create index.ts file in new project
	utils.Create_File("index.ts", generated.File__index)

	// install cors, dotenv, express, nodemon, ts-node
	cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node", "firebase-admin")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// install dev deps: cors types, express types, prisma
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "@types/cors", "@types/express", "prisma")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make app.ts
	utils.Create_File("app.ts", generated.File__firebaseAuthApp)

	// make firebase service account key file
	utils.Create_File("serviceAccountKey.json", generated.File__serviceAccountKey)

	// make dockerfile
	utils.Revise_File("docker-compose.yml", generated.File__docker, docker_port)

	fmt.Println("Composing Docker container...")

	// get docker up
	cmd_docker := utils.BoundCommand("docker", "compose", "up", "-d")
	if err := cmd_docker.Run(); err != nil {
		fmt.Println(err)
		return
	}

	utils.Work_wrapper(func() {
		// initialize primsa
		cmd_prisma := utils.BoundCommand("npx", "prisma", "init", "--datasource-provider", "postgreSQL")
		if err := cmd_prisma.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// replace the .env file
		err := os.Remove(".env")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Revise_File(".env", generated.File__firebaseEnv, docker_port)

		// replace the gitignore file
		err = os.Remove(".gitignore")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File(".gitignore", generated.File__firebaseGitignore)

		// cd into prisma
		err = os.Chdir("prisma")
		if err != nil {
			fmt.Println(err)
			return
		}

		// remove the schema and create a new one
		err = os.Remove("schema.prisma")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File("schema.prisma", generated.File__firebasePrismaSchema)

		// cd out of prisma
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

	}, "Setting up Prisma...")()

	// run a db migration
	cmd_migration := utils.BoundCommand("npx", "prisma", "migrate", "dev")
	if err := cmd_migration.Run(); err != nil {
		fmt.Println(err)
		return
	}

	utils.Work_wrapper(func() {
		// make utils folder, cd into it
		utils.Mkdir_chdir("utils")

		// create client.ts
		utils.Create_File("client.ts", generated.File__client)

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
