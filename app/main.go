package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var commandList = []string{"echo", "test", "type", "exit"}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		validCommand, arguments := validateCommand(command)

		runCommand(validCommand, arguments)
	}
}

// execute the commands
func runCommand(command string, arguments []string) {

	if command == "echo" {
		fmt.Println(strings.Join(arguments, " "))
	}

	if command == "type" {

		if validation(arguments[0]) {

			if checkCommandType(arguments[0]) {
				fmt.Println(arguments[0], "is a shell builtin")
			} else {
				path, err := locateExecutableFiles(arguments[0])

				// fmt.Println(arguments[0], "is", path)

				if err != nil {
					// fmt.Print(err)
					fmt.Print(arguments[0], ": not found\n")
				}

			}

			// exec.LookPath("pwd")
			// fmt.Println(arguments[0], "is a shell builtin")

			// path, err := locateExecutableFiles(arguments[0])

			// fmt.Print(arguments[0], " is ", path)

			// if err != nil {
			// 	fmt.Print(err)
			// 	// fmt.Print(arguments[0], ": not found\n")
			// }
			// fmt.Print(locateExecutableFiles(arguments[0]))
		} else {

			fmt.Print(arguments[0], ": not found\n")

		}

	}

	// fmt.Print("the command -> " , command , "ARGS -> " , arguments)
}

// check if the command is valid
func validateCommand(input string) (string, []string) {

	word := strings.Fields(input)
	command := word[0]
	arguments := word[1:]

	// check if the command exists
	for i := range len(commandList) {

		if command == commandList[i] {
			return command, arguments
		}

	}

	fmt.Print(command, ": command not found\n")
	return "", nil

}

func validation(input string) bool {
	found := false
	for i := range len(commandList) {
		if input == commandList[i] {
			found = true
		}
	}

	path, err := locateExecutableFiles(input)
	if err != nil {
		found = true
	}

	if path != "" {
		found = true
	}

	return found
}

// checking the command type , work only after the validation !
func checkCommandType(target string) bool {
	// fmt.Print(target)
	found := false
	for i := range len(commandList) {
		// fmt.Print(target)
		if target == commandList[i] {
			// fmt.Print(target)
			found = true
		}
	}

	return found
}

func locateExecutableFiles(target string) (string, error) {
	path, err := exec.LookPath(target)
	// fmt.Print("PATH --> ", path, "ERR --> ", err)
	return path, err
}

// func checkLocateExecutableFiles(target string) bool {
// 	_, err := exec.LookPath(target)
// 	if err != nil {
// 		return true
// 	}
// 	return false

// }

//checking the command base of the first word
//parsing the command and getting the value

//TODO check for empty string
// create a list for valid commands

// $ type echo
// echo is a shell builtin
// $ type exit
// exit is a shell builtin
// $ type type
// type is a shell builtin
// $ type invalid_command
// invalid_command: not found

//exec.LookPath("ls")
