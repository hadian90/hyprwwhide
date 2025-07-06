package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var RevealAllCmd = &cli.Command{
	Name:    "reveal-all",
	Aliases: []string{"ra"},
	Usage:   "Reveal all hidden windows on the current workspace",
	Action: func(c *cli.Context) error {
		fmt.Println("Revealing all hidden windows on the current workspace.")
		return nil
	},
}
