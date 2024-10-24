// Package vdf provides utilities to read metadata files (e.g., info.txt and changelog.txt)
// for mods and workshop items, particularly for Space Engineers Workshop content.
package vdf

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

// Readme reads the `info.txt` file located at the specified mod path and returns the title and description.
//
// The `info.txt` file is expected to follow this format:
// - Line 1: Title
// - Line 2+: Description
//
// If the file is found and read successfully, the function returns true, along with the title and description.
// If the file is not found or any error occurs, it returns false with empty strings.
//
// Parameters:
// - modPath: The path to the mod directory containing the `info.txt` file.
//
// Returns:
// - success: A boolean indicating if the file was successfully read.
// - title: The title (first line) of the `info.txt` file.
// - description: The description (remaining lines) of the `info.txt` file.
func Readme(modPath string) (bool, string, string) {
	infoFilePath := filepath.Join(modPath, "info.txt")
	debug.Print("Fetching Workshop Info at", infoFilePath)

	// Attempt to open the info.txt file
	file, err := os.Open(infoFilePath)
	if err != nil {
		return false, "", ""
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return false, "", ""
	}

	// Return the title (first line) and description (remaining lines)
	if len(lines) > 0 {
		return true, lines[0], strings.Join(lines[1:], "\n")
	}
	return false, "", ""
}

// Changelog reads the `changelog.txt` file located at the specified mod path and returns its content as a string.
//
// If the file is found and read successfully, the function returns true and the full content of the changelog.
// If the file is not found or any error occurs, it returns false with an empty string.
//
// Parameters:
// - modPath: The path to the mod directory containing the `changelog.txt` file.
//
// Returns:
// - success: A boolean indicating if the file was successfully read.
// - changelog: The full content of the `changelog.txt` file as a single string.
func Changelog(modPath string) (bool, string) {
	changelogFilePath := filepath.Join(modPath, "changelog.txt")

	// Attempt to open the changelog.txt file
	file, err := os.Open(changelogFilePath)
	if err != nil {
		return false, ""
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return false, ""
	}

	// Return the full content of the changelog as a single string
	return true, strings.Join(lines, "\n")
}
