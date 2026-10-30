package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func runCommand(command string , arguments []string) {

	fmt.Print("the command -> " , command , "ARGS -> " , arguments)
}















func validateCommand(input string) (string , []string) {

	var commandList = []string{"echo"}

	word := strings.Fields(input)
	command := word[0]
	arguments := word[1:]


	// check if the command exists
	for i := range len(commandList) {

		if command == commandList[i] {

			return command, arguments

		} else {
			fmt.Print(command, ": command not found\n")
		}
	}

	return "" , nil
}

//checking the command base of the first word
//parsing the command and getting the value

//TODO check for empty string
// create a list for valid commands
