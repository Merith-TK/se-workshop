package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Merith-TK/se-workshop/workshop/blueprint"
	"github.com/Merith-TK/se-workshop/workshop/vdf"
	"github.com/Merith-TK/utils/debug"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "workshop",
	Short: "A tool for managing Space Engineers blueprints",
	Long:  `A CLI tool to manage Space Engineers blueprints and interact with the Steam Workshop.`,
}
var (
	blueprintsDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\Blueprints\\local\\"
	commands      = map[string]string{
		"build-vdf":  "Build a VDF file from workshop ID.",
		"download":   "Download a workshop item.",
		"fix-bp":     "Fix a blueprint with a missing workshop ID.",
		"folder":     "Print the blueprints folder path.",
		"get-id":     "Get the workshop ID for a blueprint.",
		"help":       "Display this help message.",
		"login":      "Login to Steam.",
		"update":     "Update blueprints.",
		"vent-steam": "Stop and start Steam.",
	}
)

func printHelp() {
	fmt.Println("Available commands:")
	// Create a slice to hold the command names
	commandNames := make([]string, 0, len(commands))
	for cmd := range commands {
		commandNames = append(commandNames, cmd)
	}
	// Sort command names alphabetically
	sort.Strings(commandNames)

	// Print commands with descriptions
	for _, cmd := range commandNames {
		fmt.Printf("  - %s: %s\n", cmd, commands[cmd])
	}
}

func main() {
	setupSteamCMD()

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
		printHelp() // Show help if no arguments are provided
		return
	}

	switch args[0] {
	case "build-vdf":
		workshopid := blueprint.WorkshopID(args[1])
		workshopvdf := vdf.BuildVDF(workshopid, args[1])
		println(workshopvdf)
	case "download":
		steamcmd("+workshop_download_item", "244850", args[1], "+quit")
	case "folder":
		println(blueprintsDir)
	case "get-id":
		workshopid := blueprint.WorkshopID(args[1])
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + workshopid)
	case "help":
		printHelp()
	case "login":
		if len(args) > 2 {
			steamcmd("+login", args[1], args[2], "+quit")
		} else {
			steamcmd("+login", args[1], "+quit")
		}
		username := flag.Arg(1)
		filePath := blueprintsDir + "username.txt"
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(username)
		if err != nil {
			panic(err)
		}
	case "update":
		updateargs := args[1:]
		update(updateargs...)
	case "vent-steam":
		if err := stopSteam(); err != nil {
			fmt.Println("Error stopping Steam:", err)
			return
		}

		if err := startSteam(); err != nil {
			fmt.Println("Error starting Steam:", err)
		}
	case "wtf":
		fmt.Println(args)
	default:
		fmt.Println("Unknown command:", args[0])
		printHelp()
	}
	println()
}
