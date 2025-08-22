package cmd

import (
	"encoding/json"
	"fmt"
	"os"

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
		&cli.BoolFlag{
			Name:  "waybar",
			Value: false,
			Usage: "Show the number of hidden windows for waybar",
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
		} else if c.Bool("waybar") {
			// return json
			class := "notEmpty"
			if len(windows) == 0 {
				class = "empty"
			}
			json.NewEncoder(os.Stdout).Encode(map[string]interface{}{"text": len(windows), "class": class})
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
