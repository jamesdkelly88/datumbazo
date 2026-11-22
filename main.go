package main

import (
	"github.com/jamesdkelly88/datumbazo/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (silently ignore errors)
	_ = godotenv.Load()

	cmd.Execute()

}
