package utils

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
)

// currently not used
func Retry_Docker() string {

	// get a random port between 1 and 65535
	random_port := strconv.Itoa(rand.Intn(65535))
	if random_port == "0" {
		Retry_Docker()
	}

	// alert the user
	fmt.Println("The docker port is in use. Trying a different port...")

	// remove the docker-compose.yml file
	err := os.Remove("docker-compose.yml")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// make a new docker-compose.yml file
	Revise_File("docker-compose.yml", generated.File__docker, []Params{{Name: "docker_port", Value: random_port}})

	// get docker up
	cmd_docker := BoundCommand("docker", "compose", "up", "-d")
	if err := cmd_docker.Run(); err != nil {
		Retry_Docker()
	}

	return random_port
}

func getRandomPort() int {
	return rand.Intn(65535)
}

func ensure_Port_Available(port string) string {
	// format the address
	address := fmt.Sprintf(":%s", port)
	// create a listener
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("Docker port", port, "is in use. Trying a different port...")

		// get a random port between 1 and 65535
		random_port := getRandomPort()
		if random_port == 0 {
			getRandomPort()
		}

		// try again
		return ensure_Port_Available(strconv.Itoa(random_port))
	}

	defer listener.Close()
	fmt.Println("Docker port", port, "is available!")
	return port // Port is available
}

func GetDockerPort() string {
	// get users docker port preference
	docker_port := Input(
		"What docker port should the backend be on? (default: 10009)",
	)
	if docker_port == "" {
		docker_port = "10009"
	} else {
		// make sure the port is a number
		docker_port_int, err := strconv.Atoi(docker_port)
		if err != nil {
			fmt.Println("Docker port must be a number between 1 and 65535")
			GetDockerPort()
		}
		// make sure the port is between 1 and 65535
		if docker_port_int < 1 || docker_port_int > 65535 {
			fmt.Println("Docker port must be a number between 1 and 65535")
			GetDockerPort()
		}
	}

	available_port := ensure_Port_Available(docker_port)

	return available_port
}
