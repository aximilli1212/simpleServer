package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if (len(args) < 2) {
		fmt.Println("Please enter arguments")
		os.Exit(1)
	}
	fmt.Println("You entered: ", args[1:])
}
