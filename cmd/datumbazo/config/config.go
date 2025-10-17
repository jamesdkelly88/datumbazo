package config

import (
	"embed"
	"github.com/jamesdkelly88/datumbazo/pkg/dbzo"
)

//go:embed favicon.ico

var Embedded embed.FS
var Settings dbzo.Settings = dbzo.NewSettings()
var Version dbzo.Version = dbzo.GetVersion(true)
