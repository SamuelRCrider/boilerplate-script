package utils

import (
	"fmt"
	"os"
	"strings"
)

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Close_and_Remove_File(f *os.File) {
	// close and remove the file
	CloseFile(f)

	os.Remove(f.Name())
}

func Create_File(name string, file_content []string) {
	// create file
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// makes sure the file closes when function finishes execution
	defer CloseFile(file)

	// loop through data and write lines
	for _, v := range file_content {
		_, err := fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/*
* Revise File Params Struct
*
* name is the name of the variable to replace
* value is the value to replace it with
 */
type Params struct {
	name  string
	value string
}

func Revise_File(name string, file_content []string, params []Params) {
	// create file
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// makes sure the file closes when function finishes execution
	defer CloseFile(file)

	// loop through params
	for i, _ := range params {
		switch params[i].name {
		case "docker_port":
			for _, v := range file_content {
				// -1 replaces all instances
				_, err := fmt.Fprintln(file, strings.Replace(v, "10009", params[i].value, -1))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		case "project_name":
			for _, v := range file_content {
				// 1 replaces the first instance
				_, err := fmt.Fprintln(file, strings.Replace(v, "postgres:", params[i].value, 1))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		default:
			fmt.Println("Error: invalid parameter")
			return
		}
	}

}
