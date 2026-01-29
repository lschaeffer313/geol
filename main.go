package main

import (
	"runtime"

	"github.com/opt-nc/geol/v2/cmd"
	"github.com/opt-nc/geol/v2/utilities"
)

// Variables injected at build time with ldflags
var (
	Commit       = "none"
	Date         = "unknown"
	BuiltBy      = "GoReleaser"
	GoVersion    = runtime.Version()
	Version      = "dev"
	PlatformOs   = runtime.GOOS
	PlatformArch = runtime.GOARCH
)

func main() {
	// Initialize version variables in utilities package
	utilities.Version = Version
	utilities.Commit = Commit
	utilities.Date = Date
	utilities.BuiltBy = BuiltBy
	utilities.GoVersion = GoVersion
	utilities.PlatformOs = PlatformOs
	utilities.PlatformArch = PlatformArch

	cmd.Execute()
}
