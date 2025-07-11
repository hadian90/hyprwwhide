package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/hadian90/hyprwwhide/models"
	"github.com/stretchr/testify/assert"
)

// TestMain is used to set up and tear down test environment
func TestMain(m *testing.M) {
	// Setup test environment
	testDir := "./hyprwwhide_test"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		// Print error and exit if we can't create the directory
		println("Failed to create test directory:", err.Error())
		os.Exit(1)
	}

	// Set the data directory for tests
	SetDataDir(testDir)

	// Run tests
	exitCode := m.Run()

	// Clean up test environment
	// os.RemoveAll(testDir)

	os.Exit(exitCode)
}

// Helper function to get test directory
func getTestDir() string {
	return GetDataDir()
}

// Helper function to create a test window
func createTestWindow(id int) *models.Window {
	return &models.Window{
		Address: "test_address",
		Title:   "Test Window",
		Class:   "TestApp",
		Workspace: models.Workspace{
			ID:      id,
			Name:    "test_workspace",
			Monitor: "test_monitor",
		},
	}
}

func TestCheckIfMainFolderExist(t *testing.T) {
	// Set up test directory
	testDir := getTestDir()
	os.RemoveAll(testDir)
	SetDataDir(testDir)

	// Test function
	err := CheckIfMainFolderExist()

	// Assertions
	assert.NoError(t, err)
	_, err = os.Stat(testDir)
	assert.NoError(t, err, "Directory should exist after function call")
}

func TestDS_SaveHiddenWindow(t *testing.T) {
	// Set up test directory
	testDir := getTestDir()
	os.RemoveAll(testDir)
	SetDataDir(testDir)

	// Create test window
	window := createTestWindow(1)

	// Test function
	err := DS_SaveHiddenWindow(window)

	// Assertions
	assert.NoError(t, err)

	// Verify file was created with correct content
	filePath := filepath.Join(testDir, "1.json")
	_, err = os.Stat(filePath)
	assert.NoError(t, err, "File should exist")

	// Read file content
	content, err := os.ReadFile(filePath)
	assert.NoError(t, err)

	var windows []models.Window
	err = json.Unmarshal(content, &windows)
	assert.NoError(t, err)

	assert.Len(t, windows, 1)
	assert.Equal(t, window.Address, windows[0].Address)
	assert.Equal(t, window.Title, windows[0].Title)
	assert.Equal(t, window.Workspace.ID, windows[0].Workspace.ID)

	// Test adding a second window
	window2 := createTestWindow(1)
	window2.Address = "second_address"

	err = DS_SaveHiddenWindow(window2)
	assert.NoError(t, err)

	// Verify both windows are in the file
	content, err = os.ReadFile(filePath)
	assert.NoError(t, err)

	err = json.Unmarshal(content, &windows)
	assert.NoError(t, err)

	assert.Len(t, windows, 2)
	assert.Equal(t, "second_address", windows[1].Address)
}

func TestDS_LoadAllHiddenWindows(t *testing.T) {
	// Set up test directory
	testDir := getTestDir()
	os.RemoveAll(testDir)
	SetDataDir(testDir)

	// Create test windows
	window1 := createTestWindow(1)
	window2 := createTestWindow(1)
	window2.Address = "second_address"

	// Save windows
	DS_SaveHiddenWindow(window1)
	DS_SaveHiddenWindow(window2)

	// Test function
	windows, err := DS_LoadAllHiddenWindows(1)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, windows, 2)
	assert.Equal(t, window1.Address, windows[0].Address)
	assert.Equal(t, window2.Address, windows[1].Address)

	// Test loading from non-existent workspace
	windows, err = DS_LoadAllHiddenWindows(999)
	assert.NoError(t, err)
	assert.Len(t, windows, 0)
}

func TestDS_LoadLatestWindow(t *testing.T) {
	// Set up test directory
	testDir := getTestDir()
	os.RemoveAll(testDir)
	SetDataDir(testDir)

	// Create test windows
	window1 := createTestWindow(1)
	window2 := createTestWindow(1)
	window2.Address = "latest_address"

	// Save windows
	DS_SaveHiddenWindow(window1)
	DS_SaveHiddenWindow(window2)

	// Test function
	latestWindow, err := DS_LoadLatestWindow(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, latestWindow)
	assert.Equal(t, "latest_address", latestWindow.Address)

	// Test loading from non-existent workspace
	latestWindow, err = DS_LoadLatestWindow(999)
	assert.Error(t, err)
	assert.Nil(t, latestWindow)
	assert.Contains(t, err.Error(), "no windows found")
}

func TestDS_DeleteHiddenWindow(t *testing.T) {
	// Set up test directory
	testDir := getTestDir()
	os.RemoveAll(testDir)
	SetDataDir(testDir)

	// Create test windows
	window1 := createTestWindow(1)
	window2 := createTestWindow(1)
	window2.Address = "second_address"

	// Save windows
	DS_SaveHiddenWindow(window1)
	DS_SaveHiddenWindow(window2)

	// Test function
	err := DS_DeleteHiddenWindow(window1)

	// Assertions
	assert.NoError(t, err)

	// Verify window was deleted
	windows, err := DS_LoadAllHiddenWindows(1)
	assert.NoError(t, err)
	assert.Len(t, windows, 1)
	assert.Equal(t, "second_address", windows[0].Address)

	// Test deleting non-existent window
	nonExistentWindow := createTestWindow(1)
	nonExistentWindow.Address = "non_existent"

	err = DS_DeleteHiddenWindow(nonExistentWindow)
	assert.NoError(t, err) // Function doesn't return error if window not found

	// Verify no changes to file
	windows, err = DS_LoadAllHiddenWindows(1)
	assert.NoError(t, err)
	assert.Len(t, windows, 1)
}
