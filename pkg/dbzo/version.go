package dbzo

import (
	"fmt"
	"runtime"
)

const (
	name          = "Datumbazo"
	major  int    = 0
	minor  int    = 0
	patch  int    = 1
	suffix string = "-alpha"
)

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Suffix string
	Number string
	Full   string
}

func GetVersion(server bool) Version {
	version := Version{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Suffix: suffix,
	}
	version.Number = versionNumber()
	version.Full = versionString(server)
	return version
}

func versionNumber() string {
	return fmt.Sprintf("%d.%d.%d%s", major, minor, patch, suffix)
}

func versionString(server bool) string {
	var role string
	if server {
		role = "Server"
	} else {
		role = "Client"
	}

	return fmt.Sprintf("%s %s %s on %s %s, running %s, compiled by %s", name, role, versionNumber(), runtime.GOARCH, runtime.GOOS, runtime.Version(), runtime.Compiler)
}
