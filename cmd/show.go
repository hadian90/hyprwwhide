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
		// fmt.Println("Showing hidden window on the current workspace.")
		activeWorkspace := utils.GetActiveWorkspace()
		if activeWorkspace == nil {
			fmt.Println("Failed to get active workspace")
			return nil
		}
		// fmt.Printf("Active workspace: %s\n", activeWorkspace.Name)

		windows, err := utils.LoadAllHiddenWindows(activeWorkspace.ID)
		if err != nil {
			fmt.Println("Failed to load hidden windows")
			return nil
		}
		if c.Bool("number") {
			fmt.Printf("%d\n", len(windows))
		} else {
			fmt.Printf("Hidden windows: %v\n", windows)
		}
		return nil
	},
}
