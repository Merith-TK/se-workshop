package main

import (
	"flag"
	"os"
)

var (
	blueprintsDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\Blueprints\\local\\"
)

func main() {
	setupSteamCMD()

	flag.Parse()

	switch flag.Arg(0) {
	case "download":
		steamcmd("+workshop_download_item", "244850", flag.Arg(1), "+quit")
	case "login":
		steamcmd("+login", flag.Arg(1), flag.Arg(2), "+quit")
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
		args := flag.Args()[1:]
		update(args...)
	case "build-vdf":
		println(buildVDF(flag.Arg(1), flag.Arg(2)))
	case "get-id":
		println(getWorkshopID(flag.Arg(1)))
	}

	println()
}
