package semod

import (
	"fmt"
	"os"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/debug"
)

var modsDir = os.Getenv("APPDATA") + "\\SpaceEngineers\\Mods\\"

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
		shared.PrintHelp("MOD: Upload command not implemented yet")
	default:
		shared.PrintHelp("MOD: Unknown command: " + command)
	}
}