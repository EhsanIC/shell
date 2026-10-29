package main

import (
	"fmt"
)

func main() {

	var name string

	fmt.Print("$ ")
	fmt.Scanln(&name)
	fmt.Print("{", name, "} : ", "command not found")

}

// show $
// get input
// show error message
