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
	Action: func(c *cli.Context) error {
		fmt.Println("Showing hidden window on the current workspace.")
		activeWorkspace := utils.GetActiveWorkspace()
		if activeWorkspace == nil {
			fmt.Println("Failed to get active workspace")
			return nil
		}
		fmt.Printf("Active workspace: %s\n", activeWorkspace.Name)
		activeWindow := utils.GetActiveWindow()
		if activeWindow == nil {
			fmt.Println("Failed to get active window")
			return nil
		}
		fmt.Printf("Active window: %s\n", activeWindow.Title)
		return nil
	},
}
