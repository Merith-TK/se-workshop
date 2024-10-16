package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

var (
	blueprintsDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\Blueprints\\local\\"
)

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

	switch args[0] {
	case "download":
		steamcmd("+workshop_download_item", "244850", args[1], "+quit")
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
	case "build-vdf":
		workshopid := getWorkshopID(args[1])
		workshopvdf := buildVDF(workshopid, args[1])
		println(workshopvdf)
	case "get-id":
		workshopid := getWorkshopID(args[1])
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + workshopid)
	case "wtf":
		fmt.Println(args)
	}

	println()
}
