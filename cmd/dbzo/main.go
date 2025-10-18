package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"syscall"

	"github.com/jamesdkelly88/datumbazo/internal/config"
)

func main() {
	cfg := config.NewSettings(false)
	fmt.Println(cfg.Version.Full)
	if cfg.Client.Password == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Print("\n")
		cfg.Client.Password = string(bytePassword)
		if err != nil || cfg.Client.Password == "" {
			panic("Password not set")
		}
	}
	fmt.Printf("Connecting to %s:%d as %s\n", cfg.Client.Hostname, cfg.Client.Port, cfg.Client.Username)

	prompt := "> "
	for {
		fmt.Print(prompt)
		// TODO: Scanln plays hell with spaces, need an alternative
		// input, err := buffer.ReadString('\n')
		// if err != nil {
		// 	fmt.Printf("%v\n", err)
		// }
		// switch input {
		// case "":
		// 	continue
		// case "quit", "exit":
		// 	os.Exit(0)
		// default:
		// 	fmt.Printf("Unknown command: %s\n", input)
		// }
	}
}
