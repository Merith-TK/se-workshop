package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

const (
	vdfTemplate = `"workshopitem"
{
 "appid"		"244850"
 "publishedfileid"	"{WORKSHOP_ID}"
 "contentfolder"	"{MODPATH}"
 "visibility"		"0"
 "previewfile"		"{MODPATH}\thumb.png"
}`
)

func buildVDF(workshopID, modPath string) string {
	absPath, err := filepath.Abs(modPath)
	if err != nil {
		return ""
	}
	newVDF := strings.ReplaceAll(vdfTemplate, "{WORKSHOP_ID}", workshopID)
	newVDF = strings.ReplaceAll(newVDF, "{MODPATH}", absPath)
	return newVDF
}

func getWorkshopID(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return "0"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var idcode string
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		idcode = "0"
	}

	for i, line := range lines {
		if strings.HasPrefix(line, "      <WorkshopId>") && strings.HasSuffix(line, "</WorkshopId>") {
			debug.Print("Found WorkshopId line")
			code := strings.TrimPrefix(strings.TrimSuffix(line, "</WorkshopId>"), "      <WorkshopId>")
			if code != "" && code != "0" {
				fmt.Println("Found WorkshopId:", code)
				idcode = code
				break
			}
		}

		if strings.HasPrefix(line, "          <Id>") && strings.HasSuffix(line, "</Id>") {
			debug.Print("Found Id line")
			if lines[i+1] == "          <ServiceName>Steam</ServiceName>" {
				debug.Print("Found Steam ServiceName line")
				code := strings.TrimPrefix(strings.TrimSuffix(line, "</Id>"), "          <Id>")

				if code != "" && code != "0" {
					idcode = code
					break
				}
			}
		}
	}
	return idcode
}
