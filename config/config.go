package config

import (
	"fmt"
	"runtime"
	"time"
)

// version and buildTime are filled in during build by the Makefile
var (
	name      = "Go Modules HTTP Proxy"
	buildTime = "N/A"
	commit    = "N/A"
)

// Set updates the Version and ReleaseDate
func Set(n, v, t string) {
	name = n
	buildTime = t
	commit = v
}

// Version returns the current version of the binary
func Version() string {
	out := commit
	if commit == "N/A" {
		out = "0000000-dev"
	}

	return fmt.Sprintf("%s/%s (%s/%s)", name, out, runtime.GOOS, runtime.GOARCH)
}

// ReleaseDate returns the time of when the binary was built
func ReleaseDate() string {
	out := buildTime
	if buildTime == "N/A" {
		out = time.Now().Local().Format("02.01.2006 15:04 (local time)")
	}

	return out
}
