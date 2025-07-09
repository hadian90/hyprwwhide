package cmd

import (
	"log"

	"github.com/hadian90/hyprwwhide/config"
	"github.com/hadian90/hyprwwhide/utils"
	"github.com/urfave/cli/v2"
)

var HideActiveCmd = &cli.Command{
	Name:    "hide-active",
	Aliases: []string{"ha"},
	Usage:   "Hide the active window",
	Action: func(c *cli.Context) error {
		log.Println("Hiding active window.")

		// Step 1: Get active window
		activeWindow := utils.GetActiveWindow()

		log.Printf("Active window: %s\n", activeWindow.Title)
		log.Printf("Active window address: %s\n", activeWindow.Address)

		// Step 2: Hide the window
		err := utils.HideWindow(activeWindow, config.SPECIAL_WS_ID)
		if err != nil {
			log.Println("Failed to hide window")
			return nil
		}

		// Step 3: Save the hidden window
		err = utils.DS_SaveHiddenWindow(activeWindow)
		if err != nil {
			log.Println("Failed to save hidden window")
			// if failed to save, reveal the window
			utils.RevealWindow(activeWindow)
			return nil
		}

		// Step 4: Signal waybar to update
		return signal_waybar()
	},
}
