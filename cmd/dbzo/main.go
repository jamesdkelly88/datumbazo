package main

import (
	"github.com/jamesdkelly88/datumbazo/internal/config"
)

var cfg config.Settings

func main() {
	loadSettings()
	login()
	startRepl()
}
