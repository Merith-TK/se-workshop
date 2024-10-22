package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	steamCMD = filepath.Join(blueprintsDir, ".steamcmd/steamcmd.exe")
)

func setupSteamCMD() error {
	if _, err := os.Stat(steamCMD); os.IsNotExist(err) {
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

		err = unzip(filepath.Join(os.Getenv("TEMP"), "steamcmd.zip"), filepath.Join(blueprintsDir, ".steamcmd"))
		if err != nil {
			return fmt.Errorf("Failed to extract steamcmd.zip: %v", err)
		}

		fmt.Println("steamcmd downloaded and extracted successfully to", steamCMD)
	}

	return nil
}

func steamcmd(args ...string) (bytes.Buffer, error) {
	usernameFile := filepath.Join(blueprintsDir, "username.txt")
	outputBuffer := bytes.Buffer{}
	if _, err := os.Stat(usernameFile); err == nil {
		content, err := os.ReadFile(usernameFile)
		if err != nil {
			return outputBuffer, fmt.Errorf("Failed to read username.txt: %v", err)
		}
		args = append([]string{"+login", string(content)}, args...)
	}
	log.Println("[SE-Workshop] Running steamcmd with args:\n", args)
	cmd := exec.Command(steamCMD, args...)
	cmd.Stdin = os.Stdin

	cmd.Stdout = io.MultiWriter(os.Stdout, &outputBuffer)
	cmd.Stderr = io.MultiWriter(os.Stderr, &outputBuffer)

	return outputBuffer, cmd.Run()
}
