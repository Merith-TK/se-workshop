package shared

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/archive"
	"github.com/Merith-TK/utils/debug"
)

// SteamcmdDir is the path where SteamCMD is installed.
var (
	SteamcmdDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\.steamcmd"
)

// StartSteamClient launches the Steam client using a command-line call.
// It opens the Steam client if installed on the system.
func StartSteamClient() error {
	cmd := exec.Command("cmd", "/C", "start", "steam://open/main")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start Steam: %v", err)
	}
	return nil
}

// StopSteamClient terminates the Steam client process using taskkill.
// It forcibly closes Steam if running.
func StopSteamClient() error {
	cmd := exec.Command("taskkill", "/F", "/IM", "steam.exe")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop Steam: %v", err)
	}
	return nil
}

// SetupSteamcmd ensures that SteamCMD is available on the system.
// If SteamCMD is not present, it downloads and extracts the SteamCMD installer.
func SetupSteamcmd() error {
	steamcmdPath := filepath.Join(SteamcmdDir, "steamcmd.exe")

	if _, err := os.Stat(steamcmdPath); os.IsNotExist(err) {
		fmt.Println("steamcmd not found, downloading...")

		// Download steamcmd.zip
		resp, err := http.Get("https://steamcdn-a.akamaihd.net/client/installer/steamcmd.zip")
		if err != nil {
			return fmt.Errorf("failed to download steamcmd.zip: %v", err)
		}
		defer resp.Body.Close()

		// Create steamcmd.zip file in the TEMP directory
		out, err := os.Create(os.Getenv("TEMP") + "/steamcmd.zip")
		if err != nil {
			return fmt.Errorf("failed to create steamcmd.zip: %v", err)
		}
		defer out.Close()

		// Copy the downloaded data to steamcmd.zip
		if _, err := io.Copy(out, resp.Body); err != nil {
			return fmt.Errorf("failed to write steamcmd.zip: %v", err)
		}

		// Extract steamcmd.zip to SteamcmdDir
		if err := archive.Unzip(filepath.Join(os.Getenv("TEMP"), "steamcmd.zip"), SteamcmdDir); err != nil {
			return fmt.Errorf("failed to extract steamcmd.zip: %v", err)
		}

		fmt.Println("steamcmd downloaded and extracted successfully to", SteamcmdDir)
	}

	return nil
}

// Steamcmd runs SteamCMD with the provided arguments and handles authentication.
// It returns a buffer containing the command's output.
func Steamcmd(args ...string) (bytes.Buffer, error) {
	SetupSteamcmd()
	debug.SetTitle("CMD")
	defer debug.ResetTitle()

	usernameFile := filepath.Join(SteamcmdDir, "username.txt")
	outputBuffer := bytes.Buffer{}

	// Append login details if username.txt exists
	if _, err := os.Stat(usernameFile); err == nil {
		content, err := os.ReadFile(usernameFile)
		if err != nil {
			return outputBuffer, fmt.Errorf("failed to read username.txt: %v", err)
		}
		args = append([]string{"+login", string(content)}, args...)
	}

	log.Println("[SEW] Running steamcmd with args:\n", args)

	// Run steamcmd with arguments
	cmd := exec.Command(SteamcmdDir+"\\steamcmd.exe", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// TODO: Parse output for new item uploads to extract workshop ID
	// cmd.Stdout = io.MultiWriter(os.Stdout, &outputBuffer)
	// cmd.Stderr = io.MultiWriter(os.Stderr, &outputBuffer)

	return outputBuffer, cmd.Run()
}

// UploadWorkshop uploads a mod or blueprint to the Steam Workshop using SteamCMD.
// It takes a path to the mod/blueprint and a workshop ID as parameters.
func UploadWorkshop(path string, workshopID string) error {
	SetupSteamcmd()
	debug.SetTitle("Upload")
	defer debug.ResetTitle()

	debug.Print("Uploading to workshop ID:", workshopID)
	debug.Print("Path:", path)

	// Process requires folder and not bp.sbc
	if strings.HasSuffix(path, ".sbc") {
		path = strings.TrimSuffix(path, "bp.sbc")
	}

	fullPath, err := filepath.Abs(path)
	if err != nil {
		debug.Print("Failed to get absolute path:", err)
		return err
	}

	// Ensure the path exists
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		debug.Print("Failed to get file info:", err)
		return err
	}

	// If the path is a directory, prepare for workshop upload
	if fileInfo.IsDir() {
		debug.Print("Uploading directory:", fullPath)

		vdfItem := vdf.VDFItem{
			ContentFolder: fullPath,
			WorkshopID:    workshopID,
		}

		// Write workshop.vdf to the content folder
		vdfPath := filepath.Join(fullPath, "workshop.vdf")
		if err := os.WriteFile(vdfPath, []byte(vdf.Build(vdfItem)), 0644); err != nil {
			debug.Print("Failed to write workshop.vdf:", err)
			return err
		}

		debug.Print("Wrote workshop.vdf to:", vdfPath)

		// Run SteamCMD to upload the workshop item
		_, err := Steamcmd("+workshop_build_item", vdfPath, "+quit")
		if err != nil {
			debug.Print("Failed to upload workshop item:", err)
			return err
		}

		workshopID := GetWorkshopID(path)
		fmt.Println("Workshop URL: https://steamcommunity.com/sharedfiles/filedetails/?id=" + workshopID)
	}

	return nil
}
