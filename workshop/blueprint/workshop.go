package blueprint

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

func WorkshopID(path string) string {
	if !strings.HasSuffix(path, ".sbc") {
		path = filepath.Join(path, "bp.sbc")
	}
	debug.Print("Getting Workshop ID for", path)

	idcode := extractWorkshopID(path)
	if idcode == "0" || idcode == "" {
		debug.Print("No Workshop ID found, attempting to fix.")
		fixWorkshopID(path)              // Try to fix the Workshop ID by checking workshop.vdf
		idcode = extractWorkshopID(path) // Try extracting the ID again
	}

	return idcode
}

func extractWorkshopID(path string) string {
	file, err := os.Open(path)
	if err != nil {
		debug.Print("Error opening file:", err)
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
		debug.Print("Error scanning file:", err)
		idcode = "0"
	}

	for i, line := range lines {
		if strings.HasPrefix(line, "      <WorkshopId>") && strings.HasSuffix(line, "</WorkshopId>") {
			code := strings.TrimPrefix(strings.TrimSuffix(line, "</WorkshopId>"), "      <WorkshopId>")
			if code != "" && code != "0" {
				idcode = code
				break
			}
		}

		if strings.HasPrefix(line, "          <Id>") && strings.HasSuffix(line, "</Id>") {
			if lines[i+1] == "          <ServiceName>Steam</ServiceName>" {
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

func fixWorkshopID(path string) {
	if !strings.HasSuffix(path, ".sbc") {
		path = filepath.Join(path, "bp.sbc")
	}
	debug.Print("Fixing Workshop ID for", path)

	// Check if workshop.vdf exists
	workshopVDFPath := filepath.Join(filepath.Dir(path), "workshop.vdf")
	if _, err := os.Stat(workshopVDFPath); os.IsNotExist(err) {
		debug.Print("workshop.vdf does not exist")
		return
	}

	// Check if bp.sbc exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		debug.Print("bp.sbc does not exist")
		return
	}

	// Check if bp.sbc has a workshop ID
	workshopID := extractWorkshopID(path)
	workshopID = strings.TrimSpace(workshopID)
	if workshopID != "0" && workshopID != "" {
		debug.Print("bp.sbc already has a workshop ID: " + workshopID)
		return
	}

	// Read workshop ID from workshop.vdf
	file, err := os.Open(workshopVDFPath)
	if err != nil {
		debug.Print("Error opening workshop.vdf:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var workshopIDFromVDF string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ` "publishedfileid"`) {
			workshopIDFromVDF = strings.Trim(line[len(` "publishedfileid"`)+1:], `"`)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		debug.Print("Error scanning workshop.vdf:", err)
		return
	}

	if workshopIDFromVDF == "" {
		debug.Print("No workshop ID found in workshop.vdf")
		return
	}

	// Read bp.sbc content
	content, err := os.ReadFile(path)
	if err != nil {
		debug.Print("Error reading bp.sbc:", err)
		return
	}

	// Replace the end of the file with the new workshop ID
	newContent := strings.Replace(string(content),
		`      <Points>0</Points>
	</ShipBlueprint>
  </ShipBlueprints>
</Definitions>`,
		`	   <Points>0</Points>
      <WorkshopIds>
		<WorkshopId>
		  <Id>`+workshopIDFromVDF+`</Id>
		  <ServiceName>Steam</ServiceName>
		</WorkshopId>
	  </WorkshopIds>
	</ShipBlueprint>
  </ShipBlueprints>
</Definitions>`, 1)

	// Write the new content back to bp.sbc
	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		debug.Print("Error writing to bp.sbc:", err)
		return
	}

	debug.Print("Successfully updated bp.sbc with workshop ID: " + workshopIDFromVDF)
}
