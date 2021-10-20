package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/StiviiK/go-modules-http-proxy/cmd"
	"github.com/StiviiK/go-modules-http-proxy/config"
	"github.com/StiviiK/go-modules-http-proxy/pkg/html"
	"github.com/urfave/cli/v2"

	// Commands
	_ "github.com/StiviiK/go-modules-http-proxy/cmd/noop"
	_ "github.com/StiviiK/go-modules-http-proxy/cmd/serve"
)

// Version is set by an LDFLAG at build time representing the git tag or commit
// for the current release
var Version = "N/A"

// BuildTime is set by an LDFLAG at build time representing the timestamp at
// the time of build
var BuildTime = "N/A"

//go:embed assets/*
var assets embed.FS

func init() {
	config.Set("Go Modules HTTP Proxy", Version, BuildTime)

	html.ParseTemplates(assets)
}

func main() {
	defer panicHandler()

	app := cli.NewApp()
	app.Name = "Go Modules HTTP Proxy"
	app.Commands = *cmd.Retrieve()

	// All non-successful output should be written to stderr
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr

	if err := app.Run(os.Args); err != nil {
		/*
			if os.Getenv("DEBUG") == "1" {
				fmt.Fprintf(os.Stderr, "%+v\n\n%s", err, err.Error())
			} else {
				fmt.Fprintln(os.Stderr, err.Error())
				fmt.Fprintln(os.Stderr, "Re-run with DEBUG=1 for more info.")
			}
		*/
		panic(err)
	}
}

func panicHandler() {
	if r := recover(); r != nil {
		if os.Getenv("DEBUG") == "1" {
			fmt.Fprintf(os.Stderr, "%s\n", config.Version())
			fmt.Fprintf(os.Stderr, "Release Date: %s\n\n", config.ReleaseDate())
			panic(r)
		} else {
			fmt.Fprintln(os.Stderr, "Something unexpected happened.")
			fmt.Fprintln(os.Stderr, "If you want to help us debug the problem, please run:")
			fmt.Fprintf(os.Stderr, "DEBUG=1 %s\n", strings.Join(os.Args, " "))
			fmt.Fprintln(os.Stderr, "and send the output to https://github.com/StiviiK/go-modules-http-proxy")
			os.Exit(2)
		}
	}
}
