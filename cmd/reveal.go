package cmd

import (
	"fmt"

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
		if activeWorkspace == nil {
			return fmt.Errorf("failed to get active workspace")
		}

		if windowID == "" || windowID == "last" {
			fmt.Printf("Revealing %s hidden window.\n", windowID)

			latestWindow, err := utils.LoadLatestWindow(activeWorkspace.ID)
			if err != nil {
				return fmt.Errorf("failed to load latest window: %w", err)
			}

			fmt.Printf("Latest window: %s\n", latestWindow.Address)
			return revealWindowHandler(latestWindow)
		} else {
			fmt.Printf("Revealing window with ID: %s\n", windowID)
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
		return fmt.Errorf("failed to reveal window %s: %w", window.Address, err)
	}
	fmt.Printf("Revealed window: %s\n", window.Address)

	// Step 2: Delete from hidden windows list
	if err := utils.DeleteHiddenWindow(window); err != nil {
		return fmt.Errorf("failed to delete hidden window %s: %w", window.Address, err)
	}

	// Step 3: Focus the window
	if err := utils.FocusWindow(window); err != nil {
		return fmt.Errorf("failed to focus window %s: %w", window.Address, err)
	}
	fmt.Printf("Focused window: %s\n", window.Address)

	// Step 4: Signal waybar to update
	if err := signal_waybar(); err != nil {
		return fmt.Errorf("failed to signal waybar: %w", err)
	}

	return nil
}
