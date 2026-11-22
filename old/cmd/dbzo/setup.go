package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/jamesdkelly88/datumbazo/internal/config"
	"golang.org/x/term"
)

func loadSettings() {
	cfg = config.NewSettings(false)
	fmt.Println(cfg.Version.Full)
}

func login() {
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
	// call /version and print response
	test, code, err := invokeAPICall("/version")

	var failure string

	if err != nil {
		// fail hard as API call errored
		failure = err.Error()
	} else if code == 401 {
		failure = "Login failed for " + cfg.Client.Username
	} else if code != 200 {
		failure = fmt.Sprintf("Api responded with %d during login, cannot continue", code)
	}
	if failure != "" {
		fmt.Println(failure)
		os.Exit(1)
	}
	// TODO: convert version endpoint to json and compare cli/server, warn if different
	fmt.Println(test)

}
