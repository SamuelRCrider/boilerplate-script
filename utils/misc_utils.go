package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
)

func BoundCommand(name string, arg ...string) *exec.Cmd {
	command := exec.Command(name, arg...)

	// bind command to terminal
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command
}

func Open(url string) {

	switch runtime.GOOS {
	case "windows":
		// cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
		cmd := BoundCommand("rundll32", "url.dll,FileProtocolHandler", url)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	case "darwin":
		cmd := BoundCommand("open", url)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd := BoundCommand("xdg-open", url)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func Select(label string, opts []string) string {
	res := ""
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func Input(label string) string {
	res := ""
	prompt := &survey.Input{
		Message: label,
	}
	survey.AskOne(prompt, &res)

	return res
}

func Work_wrapper(wrapped func(), suffix string) func() {
	return func() {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Color("magenta")
		s.Suffix = " " + suffix
		s.Start()

		// call internal function
		wrapped()

		s.Stop()
	}

}

func Mkdir_chdir(dir string) {
	// create dir
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// cd into dir
	err = os.Chdir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
}
