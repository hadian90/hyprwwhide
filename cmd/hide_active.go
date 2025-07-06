package cmd

import (
	"fmt"

	"github.com/hadian90/hyprwwhide/config"
	"github.com/hadian90/hyprwwhide/utils"
	"github.com/urfave/cli/v2"
)

var HideActiveCmd = &cli.Command{
	Name:    "hide-active",
	Aliases: []string{"ha"},
	Usage:   "Hide the active window",
	Action: func(c *cli.Context) error {
		fmt.Println("Hiding active window.")
		activeWindow := utils.GetActiveWindow()
		if activeWindow == nil {
			return nil
		}
		fmt.Printf("Active window: %s\n", activeWindow.Title)
		fmt.Printf("Active window address: %s\n", activeWindow.Address)
		err := utils.HideWindow(activeWindow, config.SPECIAL_WS_ID)
		if err != nil {
			fmt.Println("Failed to hide window")
			return nil
		}
		err = utils.SaveHiddenWindow(activeWindow)
		if err != nil {
			fmt.Println("Failed to save hidden window")
			return nil
		}
		return nil
	},
}
