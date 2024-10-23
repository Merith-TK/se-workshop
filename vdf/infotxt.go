package vdf

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

/*
	Info.TXT
	Line 1: Title
	Line 2+: Description

*/

//TODO: Expand this to include more information

func Readme(modPath string) (bool, string, string) {
	infoFilePath := filepath.Join(modPath, "info.txt")
	debug.Print("Fetching Workshop Info at", infoFilePath)
	file, err := os.Open(infoFilePath)
	if err != nil {
		return false, "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return false, "", ""
	}

	if len(lines) > 0 {
		return true, lines[0], strings.Join(lines[1:], "\n")
	}
	return false, "", ""
}

func Changelog(modPath string) (bool, string) {
	changelogFilePath := filepath.Join(modPath, "changelog.txt")
	file, err := os.Open(changelogFilePath)
	if err != nil {
		return false, ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return false, ""
	}

	return true, strings.Join(lines, "\n")
}
