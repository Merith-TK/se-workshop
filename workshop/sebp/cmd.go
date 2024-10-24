package sebp

import (
	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/utils/debug"
)

var Directory = shared.SEDir + "\\Blueprints\\local\\"

func HandleCommand(args []string) {
	debug.SetTitle("Handling Command")
	defer debug.ResetTitle()

	command := args[0]
	args = args[1:]
	switch command {
	case "folder":
		println(Directory)
	default:
		shared.PrintHelp("BP: Unknown command: " + command)
	}
}
