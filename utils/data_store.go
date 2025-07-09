package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hadian90/hyprwwhide/models"
)

func CheckIfMainFolderExist() error {
	if _, err := os.Stat("/tmp/hyprwwhide"); err != nil {
		// create folder
		os.Mkdir("/tmp/hyprwwhide", 0755)
	}
	return nil
}

func DS_SaveHiddenWindow(window *models.Window) error {
	filePath := fmt.Sprintf("/tmp/hyprwwhide/%d.json", window.Workspace.ID)
	var windows []models.Window

	// Check if the file already exists and read existing data
	if _, err := os.Stat(filePath); err == nil {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(fileContent, &windows)
		if err != nil {
			return err
		}
	}

	// Append the new window to the existing windows
	windows = append(windows, *window)

	// Write the updated list of windows back to the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(windows)
	if err != nil {
		return err
	}

	fmt.Printf("Window saved to %s\n", filePath)
	return nil
}

func DS_LoadAllHiddenWindows(workspaceID int) ([]models.Window, error) {
	filePath := fmt.Sprintf("/tmp/hyprwwhide/%d.json", workspaceID)
	var windows []models.Window

	if _, err := os.Stat(filePath); err == nil {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(fileContent, &windows)
		if err != nil {
			return nil, err
		}
	}

	return windows, nil
}

func DS_LoadLatestWindow(workspaceID int) (*models.Window, error) {
	filePath := fmt.Sprintf("/tmp/hyprwwhide/%d.json", workspaceID)
	var windows []models.Window

	if _, err := os.Stat(filePath); err == nil {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(fileContent, &windows)
		if err != nil {
			return nil, err
		}
	}

	if len(windows) > 0 {
		return &windows[len(windows)-1], nil
	}

	return nil, fmt.Errorf("no windows found for workspace ID %d", workspaceID)
}

func DS_DeleteHiddenWindow(window *models.Window) error {
	filePath := fmt.Sprintf("/tmp/hyprwwhide/%d.json", window.Workspace.ID)
	var windows []models.Window

	// Check if the file already exists and read existing data
	if _, err := os.Stat(filePath); err == nil {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(fileContent, &windows)
		if err != nil {
			return err
		}
	}

	// Remove the window from the list
	for i, w := range windows {
		if w.Address == window.Address {
			windows = append(windows[:i], windows[i+1:]...)
			break
		}
	}

	// Write the updated list of windows back to the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(windows)
	if err != nil {
		return err
	}

	fmt.Printf("Window deleted from %s\n", filePath)
	return nil
}
