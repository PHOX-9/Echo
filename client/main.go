package main

import (
	"flag"
	"fmt"
)

func main() {
	serverURL := flag.String("server", "localhost:8080", "Server URL")

	flag.Parse()

	fmt.Println("Connecting to server...")

	username := getUsername()

	if username == "" {
		fmt.Println("Username cannot be empty")
		return
	}

	err := connectToEchoServer(*serverURL, username)
	if err != nil {
		fmt.Println(err)
		return
	}
}
