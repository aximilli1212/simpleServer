package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}
	//192051972599 - Token

	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("Value passed is not valid: %s \n", err)
		fmt.Println("You entered: ", args[1:])
		os.Exit(1)
	}
}
