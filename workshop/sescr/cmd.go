package sescr

import (
	"fmt"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/steam"
	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/se-workshop/workshop/semod"
	"github.com/Merith-TK/utils/debug"
)

var modsDir = shared.SEDir + "\\IngameScripts\\local\\"

func HandleCommand(args []string) {
	debug.SetTitle("Handling Command")
	defer debug.ResetTitle()

	command := args[0]
	args = args[1:]
	switch command {

	case "folder":
		println(modsDir)
	case "get-id", "getid", "get", "id":
		if len(args) == 0 {
			shared.PWD()
			return
		}
		workshopid := semod.WorkshopID(args[0])
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + workshopid)
	case "vdf":
		workshopid := semod.WorkshopID(args[0])
		workshopItem := vdf.VDFItem{
			WorkshopID:    workshopid,
			ContentFolder: args[0],
		}
		workshopvdf := vdf.Build(workshopItem)
		println(workshopvdf)
	case "upload", "update":
		fullpath := args[0]
		args = args[1:]
		steam.Upload(semod.WorkshopID(fullpath), args...)
	default:
		shared.PrintHelp("MOD: Unknown command: " + command)
	}
}
