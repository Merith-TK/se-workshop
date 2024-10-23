package sescr

import (
	"fmt"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf"
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
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + shared.GetWorkshopID(args[0]))
	case "vdf":
		workshopid := shared.GetWorkshopID(args[0])
		workshopItem := vdf.VDFItem{
			WorkshopID:    workshopid,
			ContentFolder: args[0],
		}
		workshopvdf := vdf.Build(workshopItem)
		println(workshopvdf)
	case "upload", "update":
		err := shared.UploadWorkshop(args[0], shared.GetWorkshopID(args[0]))
		if err != nil {
			fmt.Println("Failed to upload blueprint: " + err.Error())
		}
	default:
		shared.PrintHelp("MOD: Unknown command: " + command)
	}
}
