package noop

import (
	"github.com/StiviiK/go-modules-http-proxy/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	command := &cli.Command{
		Name:  "noop",
		Usage: "no operation",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	cmd.Register(command)
}
