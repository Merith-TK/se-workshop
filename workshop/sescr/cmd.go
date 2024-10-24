package sescr

import (
	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/utils/debug"
)

// HandleCommand processes the provided command arguments related to Space Engineers scripts.
//
// Supported commands:
//   - "folder": Prints the directory path where local scripts are stored.
//   - If the command is unrecognized, it prints an error message.
//
// Parameters:
//   - args: A slice of strings representing the command and its arguments.
//
// Usage:
//
//	sescr.HandleCommand([]string{"folder"})
func HandleCommand(args []string) {
	debug.SetTitle("Handling Command")
	defer debug.ResetTitle()

	// Ensure that at least one argument (the command) is provided.
	if len(args) == 0 {
		shared.PrintHelp("MOD: No command provided")
		return
	}

	// Extract the command and any additional arguments.
	command := args[0]
	args = args[1:]

	// Handle supported commands.
	switch command {
	case "folder":
		// Print the mods directory path.
		println(shared.Constants.Dir.Script)
	default:
		// Print an error message for unknown commands.
		shared.PrintHelp("MOD: Unknown command: " + command)
	}
}
