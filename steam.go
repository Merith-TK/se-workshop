package main

import (
	"fmt"
	"os/exec"
)

// stopSteam stops the Steam client without closing games
func stopSteam() error {
	cmd := exec.Command("taskkill", "/F", "/IM", "steam.exe")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to stop Steam: %v", err)
	}
	return nil
}

// startSteam starts the Steam client
func startSteam() error {
	cmd := exec.Command("cmd", "/C", "start", "steam://open/main")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start Steam: %v", err)
	}
	return nil
}
