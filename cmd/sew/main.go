package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/steam"
	"github.com/Merith-TK/se-workshop/workshop/blueprint"
	"github.com/Merith-TK/utils/debug"
)

func main() {
	steam.Setup()

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
		blueprint.HandleCommand(args[1:])
	case "cmd":
		steam.CMD(args[1:]...)

	case "help":
		shared.PrintHelp("")
	case "login":
		if len(args) > 2 {
			steam.CMD("+login", args[1], args[2], "+quit")
		} else {
			steam.CMD("+login", args[1], "+quit")
		}
		username := flag.Arg(1)
		filePath := filepath.Join(steam.SteamCMD, "username.txt")
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
		if err := steam.StopClient(); err != nil {
			fmt.Println("Error stopping Steam:", err)
			return
		}

		if err := steam.StartClient(); err != nil {
			fmt.Println("Error starting Steam:", err)
		}

	case "wtf":
		fmt.Println(args)
	default:
		shared.PrintHelp("Unknown command: " + args[0])
	}
}
