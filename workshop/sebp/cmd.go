package sebp

import (
	"fmt"
	"os"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/debug"
)

var blueprintsDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\Blueprints\\local\\"

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
	case "upload":
		shared.PrintHelp("BP: Upload command not implemented yet")
	default:
		shared.PrintHelp("BP: Unknown command: " + command)
	}
}
