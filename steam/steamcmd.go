package steam

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
	SteamCMD = os.Getenv("APPDATA") + "\\SpaceEngineers\\.steamcmd"
)

func Setup() error {
	if _, err := os.Stat(filepath.Join(SteamCMD, "steamcmd.exe")); os.IsNotExist(err) {
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

		err = archive.Unzip(filepath.Join(os.Getenv("TEMP"), "steamcmd.zip"), SteamCMD)
		if err != nil {
			return fmt.Errorf("Failed to extract steamcmd.zip: %v", err)
		}

		fmt.Println("steamcmd downloaded and extracted successfully to", SteamCMD)
	}

	return nil
}

func CMD(args ...string) (bytes.Buffer, error) {
	debug.SetTitle("CMD")
	defer debug.ResetTitle()

	usernameFile := filepath.Join(SteamCMD, "username.txt")
	outputBuffer := bytes.Buffer{}
	if _, err := os.Stat(usernameFile); err == nil {
		content, err := os.ReadFile(usernameFile)
		if err != nil {
			return outputBuffer, fmt.Errorf("Failed to read username.txt: %v", err)
		}
		args = append([]string{"+login", string(content)}, args...)
	}
	log.Println("[SEW] Running steamcmd with args:\n", args)
	cmd := exec.Command(SteamCMD+"\\steamcmd.exe", args...)
	cmd.Stdin = os.Stdin

	// TODO: Parse output incase of upload new item, and get the workshop id
	// cmd.Stdout = io.MultiWriter(os.Stdout, &outputBuffer)
	// cmd.Stderr = io.MultiWriter(os.Stderr, &outputBuffer)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return outputBuffer, cmd.Run()
}

func Upload(workshopid string, arg string) error {
	debug.SetTitle("Upload")
	defer debug.ResetTitle()

	debug.Print("Uploading to workshop ID:", workshopid)
	debug.Print("Args:", arg)

	debug.Print("Starting for loop")

	debug.Print("Uploading", arg)
	fullpath, abserr := filepath.Abs(arg)
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
		CMD("+workshop_build_item", vdfpath, "+quit")

	}
	return nil
}
