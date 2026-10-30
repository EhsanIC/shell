package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandList = []string{"echo" , "test" , "type"	}

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

		validCommand , arguments := validateCommand(command)

		runCommand(validCommand , arguments)
	}
}


// execute the commands
func runCommand(command string , arguments []string) {

	if command == "echo" {
		fmt.Println(strings.Join(arguments, " "))
	}

	if command == "type" {

		found := false

		for i := range len(commandList) {

			if arguments[0] == commandList[i] {
				found = true
				fmt.Println( arguments[0] , "is a shell builtin")
			}

		}

		if !found {
			fmt.Print(arguments[0], ": command not found\n")
		}

		// _ , args := validateCommand(command)
		// fmt.Print(args[0])
		// if arguments == nil {
		// 	fmt.Print("test is fine")
		// }

		// fmt.Println( arguments[0] , "is a shell builtin")
	}

	// fmt.Print("the command -> " , command , "ARGS -> " , arguments)
}














// check if the command is valid
func validateCommand(input string) (string , []string) {

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
	return "" , nil

}

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
