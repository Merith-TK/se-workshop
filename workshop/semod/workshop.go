package semod

import (
	"encoding/xml"
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
	debug.SetTitle("Getting Workshop ID")
	defer debug.ResetTitle()
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
	debug.SetTitle("Extracting Workshop ID")
	defer debug.ResetTitle()
	if !strings.HasSuffix(path, ".sbmi") {
		path = filepath.Join(path, "modinfo.sbmi")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		debug.Print("Error opening file:", err)
		return "0"
	}

	var metadata Metadata
	err = xml.Unmarshal(file, &metadata)
	if err != nil {
		debug.Print("Error unmarshalling XML:", err)
		return "0"
	}

	if metadata.WorkshopId != "" {
		return metadata.WorkshopId
	} else {

		for _, item := range metadata.WorkshopIds {
			if item.WorkshopId.ServiceName == "Steam" {
				return item.WorkshopId.ID
			}
		}
	}

	debug.Print("No Workshop ID found")
	return "0"
}

// fixWorkshopID attempts to update the .sbc file by reading the WorkshopID from workshop.vdf
func fixWorkshopID(path string) string {
	debug.SetTitle("Fixing Workshop ID")
	defer debug.ResetTitle()

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

	// parse the XML
	var metadata Metadata
	err = xml.Unmarshal(content, &metadata)
	if err != nil {
		debug.Print("Error unmarshalling XML:", err)
		return "0"
	}

	// Update the Workshop ID
	metadata.WorkshopId = vdfItem.WorkshopID
	found := false
	for i, item := range metadata.WorkshopIds {
		if item.WorkshopId.ServiceName == "Steam" {
			metadata.WorkshopIds[i].WorkshopId.ID = vdfItem.WorkshopID
			found = true
			break
		}
	}

	if !found {
		metadata.WorkshopIds = append(metadata.WorkshopIds, shared.WorkshopIDItem{
			WorkshopId: struct {
				Text        string `xml:",chardata"`
				ID          string `xml:"Id,omitempty"`
				ServiceName string `xml:"ServiceName,omitempty"`
			}{
				ServiceName: "Steam",
				ID:          vdfItem.WorkshopID,
			},
		})
	}

	// Marshal the updated metadata back to XML
	updatedContent, err := xml.MarshalIndent(metadata, "", "  ")
	if err != nil {
		debug.Print("Error marshalling XML:", err)
		return "0"
	}

	// Write the updated content back to the file
	err = os.WriteFile(path+".udated.sbc", updatedContent, 0644)
	if err != nil {
		debug.Print("Error writing file:", err)
		return "0"
	}

	return vdfItem.WorkshopID
}
