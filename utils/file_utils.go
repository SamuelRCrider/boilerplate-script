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
	Name  string
	Value string
}

func Revise_File(name string, file_content []string, params []Params) {
	// Create file
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Makes sure the file closes when function finishes execution
	defer CloseFile(file)

	// Create a copy of file_content to modify
	modified_content := make([]string, len(file_content))
	copy(modified_content, file_content)

	// Loop through params
	for _, param := range params {
		switch param.Name {
		case "docker_port":
			for i, v := range modified_content {
				// -1 replaces all instances
				modified_content[i] = strings.Replace(v, "10009", param.Value, -1)
			}
		case "project_name":
			for i, v := range modified_content {
				// we only want to replace the first instance in the file, so we break the loop after the first replacement
				newStr := strings.Replace(v, "postgres:", param.Value+":", -1)
				if v != newStr {
					modified_content[i] = newStr
					break
				}

			}
		/* add more cases here */
		default:
			fmt.Println("Error: invalid parameter")
			return
		}
	}

	// Write the modified content to the file
	for _, line := range modified_content {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
