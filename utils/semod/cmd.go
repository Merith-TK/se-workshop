package semod

import (
	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/utils/debug"
)

// HandleCommand processes the provided command arguments related to Space Engineers mods.
//
// Supported commands:
//   - "folder": Prints the directory path where local blueprints are stored.
//   - If the command is unrecognized, it prints an error message.
//
// Parameters:
//   - args: A slice of strings representing the command and its arguments.
//
// Usage:
//
//	semod.HandleCommand([]string{"folder"})
func HandleCommand(args []string) {
	debug.SetTitle("Handling Command")
	defer debug.ResetTitle()

	// Ensure that at least one argument (the command) is provided.
	if len(args) == 0 {
		shared.PrintHelp("BP: No command provided")
		return
	}

	// Extract the command and any additional arguments.
	command := args[0]
	args = args[1:]

	// Handle supported commands.
	switch command {
	case "folder":
		// Print the blueprints directory path.
		println(shared.Constants.Dir.Mod)
	default:
		// Print an error message for unknown commands.
		shared.PrintHelp("BP: Unknown command: " + command)
	}
}
