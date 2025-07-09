package cmd

import (
	"fmt"
	"log"

	"github.com/hadian90/hyprwwhide/models"
	"github.com/hadian90/hyprwwhide/utils"
	"github.com/urfave/cli/v2"
)

var RevealCmd = &cli.Command{
	Name:    "reveal",
	Aliases: []string{"r"},
	Usage:   "Reveal the targeted hidden window to existing workspace (use 'last' for the last hidden window)",
	Action: func(c *cli.Context) error {
		windowID := c.Args().First()
		activeWorkspace := utils.GetActiveWorkspace()

		if windowID == "" || windowID == "last" {
			log.Printf("Revealing %s hidden window.\n", windowID)

			latestWindow, err := utils.DS_LoadLatestWindow(activeWorkspace.ID)
			if err != nil {
				return fmt.Errorf("failed to load latest window: %w", err)
			}

			log.Printf("Latest window: %s\n", latestWindow.Address)
			return revealWindowHandler(latestWindow)
		} else {
			log.Printf("Revealing window with ID: %s\n", windowID)
			// Include workspace information for the window
			loadWindow := models.Window{
				Address:   windowID,
				Workspace: *activeWorkspace,
			}
			return revealWindowHandler(&loadWindow)
		}
	},
}

func revealWindowHandler(window *models.Window) error {
	// Step 1: Reveal the window
	if err := utils.RevealWindow(window); err != nil {
		return nil
	}
	log.Printf("Revealed window: %s\n", window.Address)

	// Step 2: Delete from hidden windows list
	if err := utils.DS_DeleteHiddenWindow(window); err != nil {
		// if failed to delete, window record are still in the hidden windows list
		// but the window is already revealed
		log.Printf("Failed to delete hidden window %s: %v", window.Address, err)
		return nil
	}

	// Step 3: Focus the window
	if err := utils.FocusWindow(window); err != nil {
		return nil
	}
	log.Printf("Focused window: %s\n", window.Address)

	// Step 4: Signal waybar to update
	return signal_waybar()
}
