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

	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/archive"
	"github.com/Merith-TK/utils/debug"
)

var (
	SteamcmdDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\.steamcmd"
)

func StartSteamClient() error {
	cmd := exec.Command("cmd", "/C", "start", "steam://open/main")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start Steam: %v", err)
	}
	return nil
}

func StopSteamClient() error {
	cmd := exec.Command("taskkill", "/F", "/IM", "steam.exe")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to stop Steam: %v", err)
	}
	return nil
}

func SetupSteamcmd() error {
	if _, err := os.Stat(filepath.Join(SteamcmdDir, "steamcmd.exe")); os.IsNotExist(err) {
		fmt.Println("steamcmd not found, downloading...")
		resp, err := http.Get("https://steamcdn-a.akamaihd.net/client/installer/steamcmd.zip")
		if err != nil {
			return fmt.Errorf("Failed to download steamcmd.zip: %v", err)
		}
		defer resp.Body.Close()

		out, err := os.Create(os.Getenv("TEMP") + "/steamcmd.zip")
		if err != nil {
			return fmt.Errorf("Failed to create steamcmd.zip: %v", err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return fmt.Errorf("Failed to write steamcmd.zip: %v", err)
		}

		err = archive.Unzip(filepath.Join(os.Getenv("TEMP"), "steamcmd.zip"), SteamcmdDir)
		if err != nil {
			return fmt.Errorf("Failed to extract steamcmd.zip: %v", err)
		}

		fmt.Println("steamcmd downloaded and extracted successfully to", SteamcmdDir)
	}

	return nil
}

func Steamcmd(args ...string) (bytes.Buffer, error) {
	SetupSteamcmd()
	debug.SetTitle("CMD")
	defer debug.ResetTitle()

	usernameFile := filepath.Join(SteamcmdDir, "username.txt")
	outputBuffer := bytes.Buffer{}
	if _, err := os.Stat(usernameFile); err == nil {
		content, err := os.ReadFile(usernameFile)
		if err != nil {
			return outputBuffer, fmt.Errorf("Failed to read username.txt: %v", err)
		}
		args = append([]string{"+login", string(content)}, args...)
	}
	log.Println("[SEW] Running steamcmd with args:\n", args)
	cmd := exec.Command(SteamcmdDir+"\\steamcmd.exe", args...)
	cmd.Stdin = os.Stdin

	// TODO: Parse output incase of upload new item, and get the workshop id
	// cmd.Stdout = io.MultiWriter(os.Stdout, &outputBuffer)
	// cmd.Stderr = io.MultiWriter(os.Stderr, &outputBuffer)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return outputBuffer, cmd.Run()
}

func UploadWorkshop(path string, workshopid string) error {
	SetupSteamcmd()
	debug.SetTitle("Upload")
	defer debug.ResetTitle()

	debug.Print("Uploading to workshop ID:", workshopid)
	debug.Print("Path:", path)

	debug.Print("Starting for loop")

	debug.Print("Uploading", path)
	fullpath, abserr := filepath.Abs(path)
	if abserr != nil {
		debug.Print("Failed to get absolute path:", abserr)
		return abserr
	}

	// Check if the path is a directory
	fileinfo, staterr := os.Stat(fullpath)
	if staterr != nil {
		debug.Print("Failed to get file info:", staterr)
		return staterr
	}

	if fileinfo.IsDir() {
		debug.Print("Uploading directory:", fullpath)
		vdfitem := vdf.VDFItem{
			ContentFolder: fullpath,
			WorkshopID:    workshopid,
		}

		vdfpath := filepath.Join(fullpath, "workshop.vdf")
		err := os.WriteFile(vdfpath, []byte(vdf.Build(vdfitem)), 0644)
		if err != nil {
			debug.Print("Failed to write workshop.vdf:", err)
			return err
		}
		debug.Print("Wrote workshop.vdf to:", vdfpath)
		Steamcmd("+workshop_build_item", vdfpath, "+quit")

	}
	return nil
}
