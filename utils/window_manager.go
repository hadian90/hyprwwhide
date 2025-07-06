package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/hadian90/hyprwwhide/models"
)

func GetActiveWorkspace() *models.Workspace {
	var workspace models.Workspace
	output, err := exec.Command("hyprctl", "activeworkspace", "-j").Output()
	if err != nil {
		fmt.Println("Failed to get active workspace")
		return nil
	}
	json.Unmarshal(output, &workspace)
	return &workspace
}

func GetActiveWindow() *models.Window {
	var window models.Window
	output, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		fmt.Println("Failed to get active window")
		return nil
	}
	json.Unmarshal(output, &window)
	return &window
}

func HideWindow(window *models.Window, workspaceId string) error {
	err := exec.Command("hyprctl", "dispatch", "movetoworkspacesilent", fmt.Sprintf("%v,address:%s", workspaceId, window.Address)).Run()
	if err != nil {
		return err
	}
	return nil
}

func RevealWindow(window *models.Window) error {
	currentWS := GetActiveWorkspace()
	err := exec.Command("hyprctl", "dispatch", "movetoworkspace", fmt.Sprintf("%v,address:%s", currentWS.ID, window.Address)).Run()
	if err != nil {
		return err
	}
	return nil
}

func FocusWindow(window *models.Window) error {
	err := exec.Command("hyprctl", "dispatch", "focuswindow", fmt.Sprintf("address:%s", window.Address)).Run()
	if err != nil {
		return err
	}
	return nil
}
