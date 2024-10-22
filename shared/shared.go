package shared

import (
	"os"

	"github.com/Merith-TK/utils/debug"
)

// fileExists checks whether a file exists and is accessible
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PrintHelp(msg string) {

	if msg != "" {
		println(msg)
	}

	println("Usage: se-workshop <command> [args...]")
	println("Commands:")
	println("  build-vdf <path> - Build a VDF file for the provided blueprint path")
	println("  folder - Print the current blueprint directory")
	println("  get-id <path> - Get the workshop ID for the provided blueprint path")

	println("This is a incomplete help message.")
}

func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		debug.Print("Error getting current directory:", err)
		return ""
	}
	return pwd
}
