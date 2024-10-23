package sebp

import (
	"fmt"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/steam"
	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/debug"
)

var blueprintsDir = shared.SEDir + "\\Blueprints\\local\\"

func HandleCommand(args []string) {
	debug.SetTitle("Handling Command")
	defer debug.ResetTitle()

	command := args[0]
	args = args[1:]
	switch command {

	case "folder":
		println(blueprintsDir)
	case "get-id", "getid", "get", "id":
		if len(args) == 0 {
			shared.PWD()
			return
		}
		workshopid := WorkshopID(args[0])
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + workshopid)
	case "vdf":
		workshopid := WorkshopID(args[0])
		workshopItem := vdf.VDFItem{
			WorkshopID:    workshopid,
			ContentFolder: args[0],
		}
		workshopvdf := vdf.Build(workshopItem)
		println(workshopvdf)
	case "upload", "update":
		err := steam.Upload(WorkshopID(args[0]), args[0])
		if err != nil {
			fmt.Println("Failed to upload blueprint: " + err.Error())
		}
	default:
		shared.PrintHelp("BP: Unknown command: " + command)
	}
}
