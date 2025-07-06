package cmd

import "os/exec"

func signal_waybar() error {
	return exec.Command("pkill", "-RTMIN+7", "waybar").Run()
}
