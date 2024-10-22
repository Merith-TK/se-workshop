package fetch

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

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
