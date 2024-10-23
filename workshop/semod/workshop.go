package semod

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/vdf" // Assuming vdf package is here
	"github.com/Merith-TK/utils/debug"
)

// WorkshopID retrieves the workshop ID for a given path, attempting to fix it if missing
func WorkshopID(path string) string {
	debug.SetTitle("Get Workshop ID")
	defer debug.ResetTitle()

	if path == "" {
		path, _ = os.Getwd()
	}

	debug.Print("Getting Workshop ID for path:", path)

	// Try to extract the workshop ID from the .sbmi file
	id := extractWorkshopID(path)
	if id == "0" || id == "" {
		debug.Print("No valid Workshop ID found in modinfo.sbmi. Attempting to fix using workshop.vdf.")
		id = fixWorkshopID(path)
	}

	return id
}

// extractWorkshopID scans the provided file for a WorkshopId tag and returns its value
func extractWorkshopID(path string) string {
	debug.SetTitle("Extract Workshop ID")
	debug.Print("Extracting Workshop ID for", path)
	defer debug.ResetTitle()

	if !strings.HasSuffix(path, ".sbmi") {
		path = filepath.Join(path, "modinfo.sbmi")
	}

	// Read the modinfo.sbmi file
	file, err := os.ReadFile(path)
	if err != nil {
		debug.Print("Error opening file:", err)
		return "0"
	}

	// Unmarshal the XML content into Metadata struct
	var metadata Metadata
	err = xml.Unmarshal(file, &metadata)
	if err != nil {
		debug.Print("Error unmarshalling XML:", err)
		return "0"
	}

	// If the WorkshopId is valid (not "0" or empty), return it
	if metadata.WorkshopId != "0" && metadata.WorkshopId != "" {
		debug.Print("Found valid Workshop ID in main tag:", metadata.WorkshopId)
		return metadata.WorkshopId
	}

	// Otherwise, check the WorkshopIds list for a Steam ID
	for _, item := range metadata.WorkshopIds {
		if item.ServiceName == "Steam" {
			debug.Print("Found valid Steam Workshop ID in list:", item.ID)
			return item.ID
		}
	}

	// If no valid Workshop ID is found, return "0"
	debug.Print("No valid Workshop ID found in modinfo.sbmi.")
	return "0"
}

// fixWorkshopID attempts to update the .sbmi file by reading the WorkshopID from workshop.vdf
func fixWorkshopID(path string) string {
	debug.SetTitle("Fix Workshop ID")
	defer debug.ResetTitle()

	path, _ = filepath.Abs(path)
	path = filepath.Clean(path)

	modinfoPath := filepath.Join(path, "modinfo.sbmi")
	workshopPath := filepath.Join(path, "workshop.vdf")

	debug.Print("Fixing Workshop ID using workshop.vdf at path:", workshopPath)

	// Ensure the workshop.vdf file exists before trying to read it
	if _, err := os.Stat(workshopPath); os.IsNotExist(err) {
		debug.Print("workshop.vdf file does not exist at path:", workshopPath)
		return "0"
	}

	// Attempt to read the Workshop ID from workshop.vdf
	vdfItem, err := vdf.Read(workshopPath)
	if err != nil || vdfItem.WorkshopID == "" {
		debug.Print("Error reading workshop.vdf or WorkshopID missing:", err)
		return "0"
	}

	// Ensure the modinfo.sbmi file exists before trying to update it
	if _, err := os.Stat(modinfoPath); os.IsNotExist(err) {
		debug.Print("modinfo.sbmi file does not exist at path:", modinfoPath)
		return "0"
	}

	// Read the modinfo.sbmi file
	content, err := os.ReadFile(modinfoPath)
	if err != nil {
		debug.Print("Error opening modinfo.sbmi file:", err)
		return "0"
	}

	// Parse the XML content
	var metadata Metadata
	err = xml.Unmarshal(content, &metadata)
	if err != nil {
		debug.Print("Error unmarshalling modinfo.sbmi XML:", err)
		return "0"
	}

	// Update the WorkshopId with the value from workshop.vdf
	metadata.WorkshopId = vdfItem.WorkshopID

	// Update the Steam-specific WorkshopId if it exists
	found := false
	for i, item := range metadata.WorkshopIds {
		if item.ServiceName == "Steam" {
			metadata.WorkshopIds[i].ID = vdfItem.WorkshopID
			found = true
			break
		}
	}

	// If no Steam WorkshopId was found, append a new one
	if !found {
		metadata.WorkshopIds = append(metadata.WorkshopIds, shared.WorkshopIDEntry{
			ServiceName: "Steam",
			ID:          vdfItem.WorkshopID,
		})
	}

	// Marshal the updated metadata back to XML
	updatedContent, err := xml.MarshalIndent(metadata, "", "  ")
	if err != nil {
		debug.Print("Error marshalling updated modinfo.sbmi XML:", err)
		return "0"
	}

	// Write the updated content back to modinfo.sbmi
	err = os.WriteFile(modinfoPath, updatedContent, 0644)
	if err != nil {
		debug.Print("Error writing updated modinfo.sbmi file:", err)
		return "0"
	}

	debug.Print("Successfully updated modinfo.sbmi with Workshop ID:", vdfItem.WorkshopID)
	return vdfItem.WorkshopID
}
