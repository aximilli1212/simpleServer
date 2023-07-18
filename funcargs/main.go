package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}
	//192051972599 - Token added

	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("Value passed is not valid: %s \n", err)
		fmt.Println("You entered: ", args[1:])
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Http status code: %d\n Body: %s\n", response.StatusCode, body)

}
