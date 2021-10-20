package cmd

import (
	"github.com/urfave/cli/v2"
)

var cmds []*cli.Command

// Register adds the given command to the global list of commands.
func Register(c *cli.Command) {
	cmds = append(cmds, c)
}

// Retrieve returns all commands
func Retrieve() *[]*cli.Command {
	return &cmds
}
