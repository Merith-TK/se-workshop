package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/se-workshop/workshop/sebp"
	"github.com/Merith-TK/se-workshop/workshop/semod"
	"github.com/Merith-TK/se-workshop/workshop/sescr"
	"github.com/Merith-TK/utils/debug"
)

func main() {
	shared.SetupSteamcmd()

	flag.Parse()
	args := []string{}
	for _, arg := range flag.Args() {
		if strings.HasSuffix(arg, "\"") && !strings.HasPrefix(arg, "\"") {
			debug.Print("Trimming quotes from:", arg)
			args = append(args, strings.TrimSuffix(arg, "\""))
		} else {
			args = append(args, arg)
		}
	}
	debug.Print("Args:", args)

	if len(args) == 0 {
		shared.PrintHelp("No Arguments") // Show help if no arguments are provided
		return
	}
	if len(args) < 2 {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		args = append(args, pwd)
	}

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
		workshopid := shared.GetWorkshopID(args[1])
		workshopItem := vdf.VDFItem{
			WorkshopID:    workshopid,
			ContentFolder: args[1],
		}
		workshopvdf := vdf.Build(workshopItem)
		fmt.Println(workshopvdf)
	case "set-id", "setid", "set":
		if len(args) < 2 {
			shared.PrintHelp("BP: set-id requires a path and a workshop ID")
			return
		}
		if !strings.HasSuffix(args[0], ".sbc") {
			args[0] = args[0] + "\\bp.sbc"
		}
		shared.SetWorkshopID(args[0], args[1])
	case "fix-contents":
		shared.SetWorkshopID(args[0], shared.GetWorkshopID(args[0]))

	case "upload", "update":
		debug.Print("Upload command detected")
		err := shared.UploadWorkshop(args[0], shared.GetWorkshopID(args[0]))
		if err != nil {
			fmt.Println("Failed to upload blueprint: " + err.Error())
		}
	case "login":
		debug.Print("Login command detected")
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
	case "vent-steam":
		debug.Print("Vent-steam command detected")
		if err := shared.StopSteamClient(); err != nil {
			fmt.Println("Error stopping Steam:", err)
			return
		}

		if err := shared.StartSteamClient(); err != nil {
			fmt.Println("Error starting Steam:", err)
		}
	default:
		debug.Print("Unknown command detected")
		shared.PrintHelp("Unknown command: " + args[0])
	}
}
