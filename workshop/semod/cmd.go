package semod

import (
	"fmt"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
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

	case "fix-contents":
		shared.SetWorkshopID(args[0], shared.GetWorkshopID(args[0]))
	case "folder":
		println(blueprintsDir)
	case "get-id", "getid", "get", "id":
		if len(args) == 0 {
			shared.PWD()
			return
		}
		if !strings.HasSuffix(args[0], ".sbmi") {
			args[0] = args[0] + "\\modinfo.sbmi"
		}
		workshopid := shared.GetWorkshopID(args[0])
		fmt.Println("https://steamcommunity.com/sharedfiles/filedetails/?id=" + workshopid)
	case "set-id", "setid", "set":
		if len(args) < 2 {
			shared.PrintHelp("BP: set-id requires a path and a workshop ID")
			return
		}
		if !strings.HasSuffix(args[0], ".sbmi") {
			args[0] = args[0] + "\\modinfo.sbmi"
		}
		shared.SetWorkshopID(args[0], args[1])
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
		shared.PrintHelp("BP: Unknown command: " + command)
	}
}
