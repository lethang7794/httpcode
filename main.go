package main

import (
	"github.com/lethang7794/httpcode/cmd"
)

// Version information - set by build flags
var (
	version = "dev"
	commit  = "N/A"
	date    = "N/A"
)

func main() {
	// Set version information in cmd package
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
