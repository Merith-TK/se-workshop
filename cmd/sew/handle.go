package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf"
)

// handleVDFCommand handles the "get-vdf" command, generating a VDF string for a given workshop item.
func handleVDFCommand(args []string) {
	workshopid := shared.GetWorkshopID(args[1])
	workshopItem := vdf.VDFItem{
		WorkshopID:    workshopid,
		ContentFolder: args[1],
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
	if !strings.HasSuffix(args[0], ".sbc") {
		args[0] = args[0] + "\\bp.sbc"
	}
	shared.SetWorkshopID(args[0], args[1])
}

// handleUploadCommand handles the "upload" or "update" commands, uploading the workshop item.
func handleUploadCommand(args []string) {
	err := shared.UploadWorkshop(args[0], shared.GetWorkshopID(args[0]))
	if err != nil {
		fmt.Println("Failed to upload blueprint: " + err.Error())
	}
}

// handleLoginCommand handles the "login" command, logging into Steamcmd and saving the username.
func handleLoginCommand(args []string) {
	if len(args) > 2 {
		shared.Steamcmd("+login", args[1], args[2], "+quit")
	} else {
		shared.Steamcmd("+login", args[1], "+quit")
	}
	username := flag.Arg(1)
	filePath := filepath.Join(shared.SteamcmdDir, "username.txt")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(username)
	if err != nil {
		panic(err)
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
