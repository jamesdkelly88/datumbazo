package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	username := os.Getenv("DBZO_USER")

	args := os.Args[1:] // skip the application name (handles quoted variables automatically)

	fmt.Printf("Hello %s\n", username)

	for _, element := range args {
		fmt.Println(element)
	}
}
