package main

import (
	"log/slog"
	"os"

	"github.com/jamesdkelly88/datumbazo/internal/logging"
	"github.com/jamesdkelly88/datumbazo/internal/tokeniser"
)

func main() {
	logging.SetupLogger(os.Stdout, slog.LevelDebug)
	query, err := os.ReadFile("./query.txt")
	if err != nil {
		panic(err)
	}
	tokeniser.Test(string(query))
}
