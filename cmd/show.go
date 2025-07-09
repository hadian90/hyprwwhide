package cmd

import (
	"fmt"

	"github.com/hadian90/hyprwwhide/utils"
	"github.com/urfave/cli/v2"
)

var ShowCmd = &cli.Command{
	Name:    "show",
	Aliases: []string{"s"},
	Usage:   "Show a hidden window on the current workspace",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "number",
			Value: false,
			Usage: "Only show the number of hidden windows",
		},
	},
	Action: func(c *cli.Context) error {
		activeWorkspace := utils.GetActiveWorkspace()

		windows, err := utils.DS_LoadAllHiddenWindows(activeWorkspace.ID)
		if err != nil {
			return fmt.Errorf("failed to load hidden windows: %w", err)
		}

		if c.Bool("number") {
			fmt.Printf("%d\n", len(windows))
		} else {
			// Improve the display of windows with more detailed information
			if len(windows) == 0 {
				fmt.Println("No hidden windows in this workspace")
			} else {
				fmt.Printf("Hidden windows in workspace %d (%s):\n",
					activeWorkspace.ID, activeWorkspace.Name)
				for i, window := range windows {
					fmt.Printf("%d. Address: %s\n", i+1, window.Address)
				}
			}
		}

		return nil
	},
}
