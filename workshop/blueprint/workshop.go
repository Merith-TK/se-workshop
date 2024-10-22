package blueprint

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf" // Assuming vdf package is here
	"github.com/Merith-TK/utils/debug"
)

// WorkshopID retrieves the workshop ID for a given path, attempting to fix it if missing
func WorkshopID(path string) string {
	if path == "" {
		path, _ = os.Getwd()
	}
	debug.Print("Getting Workshop ID for", path)

	// Try to extract the workshop ID from the .sbc file
	id := extractWorkshopID(path)
	if id == "0" || id == "" {
		debug.Print("No Workshop ID found, attempting to fix.")
		id = fixWorkshopID(path)
	}

	return id
}

// extractWorkshopID scans the provided file for a WorkshopId tag and returns its value
func extractWorkshopID(path string) string {
	if !strings.HasSuffix(path, ".sbc") {
		path = filepath.Join(path, "bp.sbc")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		debug.Print("Error opening file:", err)
		return "0"
	}

	fileStr := string(file)

	// Find the Workshop ID tag
	if i := strings.Index(fileStr, "<WorkshopId>"); i != -1 {
		id := fileStr[i+len("<WorkshopId>") : strings.Index(fileStr[i:], "</WorkshopId>")+i]
		return id
	}

	if i := strings.Index(fileStr, "<WorkshopIds>"); i != -1 {
		if j := strings.Index(fileStr[i:], "<Id>"); j != -1 {
			id := fileStr[i+j+len("<Id>") : strings.Index(fileStr[i+j:], "</Id>")+i+j]
			return id
		}
	}

	return "0"
}

// fixWorkshopID attempts to update the .sbc file by reading the WorkshopID from workshop.vdf
func fixWorkshopID(path string) string {
	if !strings.HasSuffix(path, ".sbc") {
		path = filepath.Join(path, "bp.sbc")
	}
	debug.Print("Fixing Workshop ID for", path)

	// Ensure the required files exist
	workshopVDFPath := filepath.Join(filepath.Dir(path), "workshop.vdf")
	if !shared.FileExists(workshopVDFPath) || !shared.FileExists(path) {
		fmt.Println("Required files not found: either bp.sbc or workshop.vdf")
		return "0"
	}

	// Attempt to read the Workshop ID from workshop.vdf
	vdfItem, err := vdf.Read(workshopVDFPath)
	if err != nil || vdfItem.WorkshopID == "" {
		debug.Print("Error reading workshop.vdf or no WorkshopID found")
		return "0"
	}

	content, err := os.ReadFile(path)
	if err != nil {
		debug.Print("Error opening file:", err)
		return "0"
	}

	newContent := strings.Replace(string(content),
		`      <Points>0</Points>
	</ShipBlueprint>
  </ShipBlueprints>
</Definitions>`,
		`      <Points>0</Points>
      <WorkshopIds>
        <WorkshopId>
          <Id>`+vdfItem.WorkshopID+`</Id>
          <ServiceName>Steam</ServiceName>
        </WorkshopId>
      </WorkshopIds>
	</ShipBlueprint>
  </ShipBlueprints>
</Definitions>`, 1)

	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		debug.Print("Error writing to file:", err)
		return "0"

	}
	return vdfItem.WorkshopID
}
