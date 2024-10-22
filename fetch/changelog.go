package fetch

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

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
