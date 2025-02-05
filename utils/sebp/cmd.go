package sebp

import (
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/utils/debug"
)

// HandleCommand processes the provided arguments and executes the corresponding blueprint-related commands.
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
//	sebp.HandleCommand([]string{"folder"})
func HandleCommand(args []string) {
	debug.SetTitle("Handling Command")
	defer debug.ResetTitle()

	// Ensure at least one argument (the command) is provided.
	if len(args) == 0 {
		shared.PrintHelp("BP: No command provided")
		return
	}

	// Extract the command and any additional arguments.
	command := args[0]

	// Handle supported commands.
	switch command {
	case "folder":
		// Print the blueprint directory path.
		println(shared.Constants.Dir.BP)
	default:
		// Print an error message for unknown commands.
		shared.PrintHelp("BP: Unknown command: " + command)
	}
}

func LocateBP(path string) (bool, string) {
	path = filepath.Clean(path)
	if strings.HasSuffix(path, ".sbc") {
		return true, path
	}
	if shared.FileExists(path + "\\bp.sbc") {
		return true, path + "\\bp.sbc"
	}
	return false, ""
}
