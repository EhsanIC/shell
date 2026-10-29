package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString("\r\n")
	
	if err != nil {
		fmt.Fprint(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	
	fmt.Println(command[:len(command)-1] + ": command not found")

}

// show $
// get input
// show error message
