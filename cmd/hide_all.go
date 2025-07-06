package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var HideAllCmd = &cli.Command{
	Name:    "hide-all",
	Aliases: []string{"hla"},
	Usage:   "Hide all windows on the current workspace",
	Action: func(c *cli.Context) error {
		fmt.Println("Hiding all windows on the current workspace.")
		return nil
	},
}
