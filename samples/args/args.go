package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // skip the application name (handles quoted variables automatically)

	//  define defaults
	port := "8080"
	name := "Default"
	version := "1"

	for _, element := range args {
		// if parameter contains =
		if strings.Contains(element, "=") {
			parts := strings.Split(element, "=")
			switch strings.ToLower(parts[0]) {
			// update defaults
			case "-p", "-port":
				port = parts[1]
			case "-n", "-name":
				name = parts[1]
			case "-v", "-version":
				version = parts[1]
			}
		} else if element == "help" || element == "-help" || element == "--help" {
			fmt.Println("help function")
			os.Exit(0)
		}
	}

	fmt.Printf("Name: %s, Port: %s, Version: %s", name, port, version)
}
