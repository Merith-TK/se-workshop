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

	"github.com/Merith-TK/utils/archive"
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

	cmd.Stdout = io.MultiWriter(os.Stdout, &outputBuffer)
	cmd.Stderr = io.MultiWriter(os.Stderr, &outputBuffer)

	return outputBuffer, cmd.Run()
}
