package main

import (
	"os"

	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/middleware"
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/router"
	"github.com/jamesdkelly88/datumbazo/internal/config"
)

func main() {
	cfg := config.NewSettings(true)
	middleware.SetupLogger(os.Stdout, cfg.LogLevel)
	router.Serve(cfg.Server.Listen, cfg)
}
