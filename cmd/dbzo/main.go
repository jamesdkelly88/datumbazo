package main

import (
	"fmt"

	"github.com/jamesdkelly88/datumbazo/pkg/dbzo"
)

var version dbzo.Version

func main() {
	version = dbzo.GetVersion(false)
	fmt.Println(version.Full)
}
