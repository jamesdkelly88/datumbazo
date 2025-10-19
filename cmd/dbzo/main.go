package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/term"

	"github.com/jamesdkelly88/datumbazo/internal/config"
)

var prompt string = "dbzo > "

var commands = map[string]interface{}{
	".clear": clearScreen,
	".exit":  endProgram,
	".help":  showHelp,
	".quit":  endProgram,
}

func main() {
	// show version
	cfg := config.NewSettings(false)
	fmt.Println(cfg.Version.Full)
	// get password (unless provided by env vars)
	if cfg.Client.Password == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Print("\n")
		cfg.Client.Password = string(bytePassword)
		if err != nil || cfg.Client.Password == "" {
			panic("Password not set")
		}
	}
	// test credentials by connecting
	fmt.Printf("Connecting to %s:%d as %s\n", cfg.Client.Hostname, cfg.Client.Port, cfg.Client.Username)
	// TODO: call /version and print response
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	showPrompt()
	for reader.Scan() {
		input := cleanInput(reader.Text())
		if strings.HasPrefix(input, "/") {
			// Call the api directly
			output, err := invokeApiCall(cfg, input)
			if err != nil {
				fmt.Printf("Api call failed: %v\n", err)
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

func showPrompt() {
	fmt.Print(prompt)
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

func invokeApiCall(cfg config.Settings, path string) (string, error) {
	url := fmt.Sprintf("http://%s:%d%s", cfg.Client.Hostname, cfg.Client.Port, path)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(cfg.Client.Username, cfg.Client.Password)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	s := string(bodyText)
	return s, nil
}
