package express_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/utils"
)

func Express_ClerkAuth(docker_port string) {
	// mkdir for backend, 0755 is the permission bits
	utils.Mkdir_chdir("backend")

	// initialize npm project
	cmd_npm := utils.BoundCommand("npm", "init")
	if err := cmd_npm.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// create index.ts file in new project
	utils.Create_File("index.ts", generated.File__index)

	// install cors, dotenv, express, nodemon, ts-node
	cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node", "@clerk/clerk-sdk-node")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// install dev deps: cors types, express types, prisma
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "@types/cors", "@types/express", "prisma", "@clerk/types")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make app.ts
	utils.Create_File("app.ts", generated.File__expressClerkApp)

	utils.Work_wrapper(func() {
		// make dockerfile
		utils.Revise_File("docker-compose.yml", generated.File__docker, docker_port)

		// get docker up
		cmd_docker := utils.BoundCommand("docker", "compose", "up", "-d")
		if err := cmd_docker.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}, "Starting Docker container...")()

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

		utils.Revise_File(".env", generated.File__expressClerkEnv, docker_port)

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

		utils.Create_File("schema.prisma", generated.File__expressClerkSchema)

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

		// create global.d.ts file
		utils.Create_File("global.d.ts", generated.File__expressClerkGlobal)

		// create client.ts file
		utils.Create_File("client.ts", generated.File__client)

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
