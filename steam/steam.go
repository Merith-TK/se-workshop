package steam

import (
	"fmt"
	"os/exec"
)

func StartClient() error {
	cmd := exec.Command("cmd", "/C", "start", "steam://open/main")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start Steam: %v", err)
	}
	return nil
}

func StopClient() error {
	cmd := exec.Command("taskkill", "/F", "/IM", "steam.exe")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to stop Steam: %v", err)
	}
	return nil
}
