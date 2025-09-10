package main

import (
	"fmt"
	"log"
	"example/greetings"
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