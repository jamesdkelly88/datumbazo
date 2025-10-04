package dbzo

// Settings contains all the configuration for the engine, grouped into smaller structures.
type Settings struct {
	Server ServerSettings
}

// ServerSettings contains the configuration for the server connection.
type ServerSettings struct {
	Port int
}

func NewSettings() Settings {
	settings := Settings{}
	settings.Server = newServerSettings()
	return settings
}

func newServerSettings() ServerSettings {
	server := ServerSettings{}
	server.Port = 8874
	return server
}
