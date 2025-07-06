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
		if windowID == "" || windowID == "last" {
			fmt.Printf("Revealing %s hidden window.\n", windowID)
			activeWorkspace := utils.GetActiveWorkspace()
			if activeWorkspace == nil {
				fmt.Println("Failed to get active workspace")
				return nil
			}
			latestWindow, err := utils.LoadLatestWindow(activeWorkspace.ID)
			if err != nil {
				fmt.Println("Failed to load latest window")
				return nil
			}
			fmt.Printf("Latest window: %s\n", latestWindow.Address)
			return revealWindowHandler(latestWindow)
		} else {
			fmt.Printf("Revealing window with ID: %s\n", windowID)
			loadWindow := models.Window{Address: windowID}
			return revealWindowHandler(&loadWindow)
		}
	},
}

func revealWindowHandler(window *models.Window) error {
	err := utils.RevealWindow(window)
	if err != nil {
		fmt.Println("Failed to reveal window")
		return nil
	}
	fmt.Printf("Revealed window: %s\n", window.Address)
	err = utils.DeleteHiddenWindow(window)
	if err != nil {
		fmt.Println("Failed to delete hidden window")
		return nil
	}
	err = utils.FocusWindow(window)
	if err != nil {
		fmt.Println("Failed to focus window")
		return nil
	}
	fmt.Printf("Focused window: %s\n", window.Address)
	return signal_waybar()
}
