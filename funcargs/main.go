package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("You entered: ", args[1:])
}
