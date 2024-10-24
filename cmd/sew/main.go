// Package main provides a command-line interface for managing Space Engineers
// blueprints, mods, scripts, and related Steam Workshop items. It supports various
// commands for downloading, uploading, setting, and managing Steam Workshop content.
package main

import (
	"flag"
	"fmt"
	"os"

	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/workshop/sebp"
	"github.com/Merith-TK/se-workshop/workshop/semod"
	"github.com/Merith-TK/se-workshop/workshop/sescr"
	"github.com/Merith-TK/utils/debug"
)

// main is the entry point of the application. It processes command-line arguments
// and executes the appropriate command based on the first argument provided.
func main() {
	// Initialize Steamcmd setup
	shared.SetupSteamcmd()

	// Parse command-line flags
	flag.Parse()

	// Process arguments, removing trailing quotes if needed
	args := parseArgs(flag.Args())

	// If no arguments are provided, show help and exit
	if len(args) == 0 {
		shared.PrintHelp("No Arguments")
		return
	}

	// Append current working directory if only one argument is provided
	if len(args) < 2 {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		args = append(args, pwd)
	}

	// Handle the appropriate command based on the first argument
	switch args[0] {
	case "bp", "blueprint", "blueprints", "schematic", "schematics", "schem":
		debug.Print("Blueprint command detected")
		sebp.HandleCommand(args[1:])

	case "mod", "mods":
		debug.Print("Mod command detected")
		semod.HandleCommand(args[1:])

	case "script", "scripts", "scr", "src":
		debug.Print("Script command detected")
		sescr.HandleCommand(args[1:])

	case "cmd":
		debug.Print("CMD command detected")
		shared.Steamcmd(args[1:]...)

	case "download", "dl":
		debug.Print("Download command detected")
		shared.Steamcmd("+workshop_download_item", "244850", args[1], "+quit")

	case "get-id", "getid", "get", "id":
		debug.Print("Get-id command detected")
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + shared.GetWorkshopID(args[1]))

	case "get-vdf", "getvdf", "vdf":
		debug.Print("Get-vdf command detected")
		handleVDFCommand(args)

	case "set-id", "setid", "set":
		debug.Print("Set-id command detected")
		handleSetIDCommand(args)

	case "fix-contents":
		debug.Print("Fix-contents command detected")
		shared.SetWorkshopID(args[0], shared.GetWorkshopID(args[0]))

	case "upload", "update":
		debug.Print("Upload command detected")
		handleUploadCommand(args)

	case "login":
		debug.Print("Login command detected")
		handleLoginCommand(args)

	case "vent-steam":
		debug.Print("Vent-steam command detected")
		handleVentSteamCommand()

	default:
		debug.Print("Unknown command detected")
		shared.PrintHelp("Unknown command: " + args[0])
	}
}

// parseArgs processes command-line arguments, removing trailing quotes where necessary.
func parseArgs(rawArgs []string) []string {
	args := []string{}
	for _, arg := range rawArgs {
		if strings.HasSuffix(arg, "\"") && !strings.HasPrefix(arg, "\"") {
			debug.Print("Trimming quotes from:", arg)
			args = append(args, strings.TrimSuffix(arg, "\""))
		} else {
			args = append(args, arg)
		}
	}
	debug.Print("Args:", args)
	return args
}
