package config

import (
	"log/slog"
	"os/user"

	flag "github.com/spf13/pflag"
)

type Settings struct {
	Client struct {
		Hostname string
		Port     int
		Username string
		Password string
	}
	Server struct {
		Listen string
	}
	Version  Version
	LogLevel slog.Level
}

func NewSettings(server bool) Settings {
	settings := Settings{}
	settings.Version = GetVersion(server)
	settings.LogLevel = slog.LevelDebug // TODO: add flag

	if server {
		serverSettings(&settings)
	} else {
		clientSettings(&settings)
	}

	flag.Parse()

	return settings
}

func serverSettings(cfg *Settings) {
	// server arguments
	flag.StringVarP(&cfg.Server.Listen, "listen", "l", ":8874", "The hostname and port number the web application listens on")
}

func clientSettings(cfg *Settings) {
	// cli arguments
	flag.StringVarP(&cfg.Client.Hostname, "host", "h", "localhost", "The server to connect to")
	flag.IntVarP(&cfg.Client.Port, "port", "p", 8874, "The port number to connect to")
	u, err := user.Current()
	var username string
	if err == nil {
		username = u.Username
	}
	flag.StringVarP(&cfg.Client.Username, "username", "U", username, "The username to login with")
	if cfg.Client.Username == "" {
		panic("Username not set")
	}
}
