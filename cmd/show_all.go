package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var ShowAllCmd = &cli.Command{
	Name:    "show-all",
	Aliases: []string{"sa"},
	Usage:   "Show all hidden windows on all workspaces",
	Action: func(c *cli.Context) error {
		fmt.Println("Showing all hidden windows on all workspaces.")
		return nil
	},
}
