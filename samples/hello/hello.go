package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0) // disable time, file, line number
	log.SetPrefix("log: ")
	message, err := greetings.Hello("James")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
