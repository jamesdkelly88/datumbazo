package main

import (
	"github.com/jamesdkelly88/datumbazo/cmd/datumbazo/router"
	"github.com/jamesdkelly88/datumbazo/internal/config"
)

func main() {
	cfg := config.NewSettings(true)
	router.Serve(cfg.Server.Listen, cfg)
}
