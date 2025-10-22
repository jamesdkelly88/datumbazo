package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var prompt string = "dbzo"

var commands = map[string]interface{}{
	".clear": clearScreen,
	".exit":  endProgram,
	".help":  showHelp,
	".quit":  endProgram,
}

func startRepl() {
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	showPrompt()
	for reader.Scan() {
		input := cleanInput(reader.Text())
		if strings.HasPrefix(input, "/") {
			// Call the api directly
			output, code, err := invokeAPICall(input)
			if err != nil {
				fmt.Printf("Api call failed: %v\n", err)
			} else if code != 200 {
				fmt.Printf("Api responded with %d\n", code)
			} else {
				fmt.Println(output)
			}
		} else if command, exists := commands[input]; exists {
			// Call a hardcoded function
			command.(func())()
		} else {
			// Handle anything else
			// TODO: handle multi-line commands until you hit a semicolon
			printUnknown(input)
		}
		showPrompt()
	}
}

// TODO: refactor these to use a writer so they can be tested
// https://www.reddit.com/r/golang/comments/kuzxtc/question_the_best_way_to_test_a_print_statement/

func showPrompt() {
	fmt.Print(prompt + " > ")
}

func showHelp() {
	fmt.Println("No help here, just us chickens!")
}

func printUnknown(text string) {
	fmt.Println("Unknown command: ", text)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func endProgram() {
	os.Exit(0)
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	return output
}
