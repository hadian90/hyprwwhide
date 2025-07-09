package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/hadian90/hyprwwhide/models"
)

func GetActiveWorkspace() *models.Workspace {
	var workspace models.Workspace
	output, err := exec.Command("hyprctl", "activeworkspace", "-j").Output()
	if err != nil {
		log.Fatal("Failed to get active workspace")
	}
	if err = json.Unmarshal(output, &workspace); err != nil {
		log.Fatal("Failed to unmarshal active workspace")
	}
	return &workspace
}

func GetActiveWindow() *models.Window {
	var window models.Window
	output, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		log.Fatal("Failed to get active window")
	}
	if err = json.Unmarshal(output, &window); err != nil {
		log.Fatal("Failed to unmarshal active window")
	}
	return &window
}

func HideWindow(window *models.Window, workspaceId string) error {
	err := exec.Command("hyprctl", "dispatch", "movetoworkspacesilent", fmt.Sprintf("%v,address:%s", workspaceId, window.Address)).Run()
	if err != nil {
		log.Printf("Failed to hide window %s: %v", window.Address, err)
		return err
	}
	return nil
}

func RevealWindow(window *models.Window) error {
	currentWS := GetActiveWorkspace()
	err := exec.Command("hyprctl", "dispatch", "movetoworkspace", fmt.Sprintf("%v,address:%s", currentWS.ID, window.Address)).Run()
	if err != nil {
		log.Printf("Failed to reveal window %s: %v", window.Address, err)
		return err
	}
	return nil
}

func FocusWindow(window *models.Window) error {
	err := exec.Command("hyprctl", "dispatch", "focuswindow", fmt.Sprintf("address:%s", window.Address)).Run()
	if err != nil {
		log.Printf("Failed to focus window %s: %v", window.Address, err)
		return err
	}
	return nil
}
