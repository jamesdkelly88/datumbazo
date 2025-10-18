package config

import (
	flag "github.com/spf13/pflag"
	"os/user"
)

type Settings struct {
	Client struct {
		Hostname string
		Port     int
		Username string
		Password string
	}
	Server struct {
		Port int
	}
	Version Version
}

func NewSettings(server bool) Settings {
	settings := Settings{}
	settings.Version = GetVersion(server)

	if server {
		// server arguments
		flag.IntVarP(&settings.Server.Port, "port", "p", 8874, "The port number the web application listens on")
	} else {
		// cli arguments
		flag.StringVarP(&settings.Client.Hostname, "host", "h", "localhost", "The server to connect to")
		flag.IntVarP(&settings.Client.Port, "port", "p", 8874, "The port number to connect to")
		u, err := user.Current()
		var username string
		if err == nil {
			username = u.Username
		}
		flag.StringVarP(&settings.Client.Username, "username", "U", username, "The username to login with")
		if settings.Client.Username == "" {
			panic("Username not set")
		}
	}

	flag.Parse()

	return settings
}
