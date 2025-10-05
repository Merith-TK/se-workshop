package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf"
)

// handleVDFCommand handles the "get-vdf" command, generating a VDF string for a given workshop item.
func handleVDFCommand(args []string) {
	workshopid := shared.GetWorkshopID(args[0])
	workshopItem := vdf.VDFItem{
		WorkshopID:    workshopid,
		ContentFolder: args[0],
	}
	workshopvdf := vdf.Build(workshopItem)
	fmt.Println(workshopvdf)
}

// handleSetIDCommand handles the "set-id" command, setting the workshop ID for a blueprint.
func handleSetIDCommand(args []string) {
	if len(args) < 2 {
		shared.PrintHelp("BP: set-id requires a path and a workshop ID")
		return
	}

	// Validate workshop ID
	if err := shared.ValidateWorkshopID(args[1]); err != nil {
		fmt.Printf("Invalid workshop ID: %v\n", err)
		return
	}

	// Sanitize and validate path
	path := shared.SanitizePath(args[0])
	if err := shared.ValidateFilePath(path); err != nil {
		fmt.Printf("Invalid file path: %v\n", err)
		return
	}

	if !strings.HasSuffix(path, ".sbc") {
		path = filepath.Join(path, shared.BlueprintFileName)
	}
	shared.SetWorkshopID(path, args[1])
}

// handleUploadCommand handles the "upload" or "update" commands, uploading the workshop item.
func handleUploadCommand(args []string) {
	// Sanitize and validate path
	path := shared.SanitizePath(args[0])
	if err := shared.ValidateFilePath(path); err != nil {
		fmt.Printf("Invalid file path: %v\n", err)
		return
	}

	err := shared.UploadWorkshop(path, shared.GetWorkshopID(path))
	if err != nil {
		fmt.Println("Failed to upload blueprint: " + err.Error())
	}
}

// handleLoginCommand handles the "login" command, logging into Steamcmd and saving the username.
func handleLoginCommand(args []string) {
	if len(args) < 1 {
		shared.PrintHelp("Login requires at least a username")
		return
	}

	// Validate username
	if err := shared.ValidateUsername(args[0]); err != nil {
		fmt.Printf("Invalid username: %v\n", err)
		return
	}

	if len(args) > 2 {
		shared.Steamcmd("+login", args[0], args[1], "+quit")
	} else {
		shared.Steamcmd("+login", args[0], "+quit")
	}

	username := args[0]
	filePath := filepath.Join(shared.SteamcmdDir, shared.UsernameFileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating username file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(username)
	if err != nil {
		fmt.Printf("Error writing username: %v\n", err)
	}
}

// handleVentSteamCommand handles the "vent-steam" command, stopping and starting the Steam client.
func handleVentSteamCommand() {
	if err := shared.StopSteamClient(); err != nil {
		fmt.Println("Error stopping Steam:", err)
		return
	}

	if err := shared.StartSteamClient(); err != nil {
		fmt.Println("Error starting Steam:", err)
	}
}
