package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/steam"
	"github.com/Merith-TK/se-workshop/workshop/sebp"
	"github.com/Merith-TK/se-workshop/workshop/semod"
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
		debug.Print("Blueprint command detected")
		sebp.HandleCommand(args[1:])
	case "mod", "mods":
		debug.Print("Mod command detected")
		semod.HandleCommand(args[1:])
	case "cmd":
		debug.Print("CMD command detected")
		steam.CMD(args[1:]...)
	case "help":
		debug.Print("Help command detected")
		shared.PrintHelp("")
	case "login":
		debug.Print("Login command detected")
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
		debug.Print("Vent-steam command detected")
		if err := steam.StopClient(); err != nil {
			fmt.Println("Error stopping Steam:", err)
			return
		}

		if err := steam.StartClient(); err != nil {
			fmt.Println("Error starting Steam:", err)
		}

	case "wtf":
		debug.Print("WTF command detected")
		fmt.Println(args)
	default:
		debug.Print("Unknown command detected")
		shared.PrintHelp("Unknown command: " + args[0])
	}
}
