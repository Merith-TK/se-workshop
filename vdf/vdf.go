package vdf

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/fetch"
	"github.com/Merith-TK/utils/debug"
)

// VDFItem represents a generic VDF item with all relevant fields
type VDFItem struct {
	AppID         string
	WorkshopID    string
	ContentFolder string
	Visibility    string
	PreviewFile   string
	Title         string
	Description   string
	ChangeNote    string
}

// BuildVDF builds the VDF file content based on the provided VDFItem struct
func Build(item VDFItem) string {
	absPath, err := filepath.Abs(item.ContentFolder)
	if err != nil {
		return ""
	}

	if item.AppID == "" {
		item.AppID = "244850"
	}
	if item.Visibility == "" {
		item.Visibility = "0"
	}
	if item.PreviewFile == "" {
		item.PreviewFile = filepath.Join(absPath, "thumb.png")
	}
	if item.WorkshopID == "" {
		return ""
	}

	// Start building the VDF content
	var vdfContent []string
	vdfContent = append(vdfContent, `"workshopitem"`)
	vdfContent = append(vdfContent, `{`)

	// Required fields
	vdfContent = append(vdfContent, ` "appid"		"`+item.AppID+`"`)
	vdfContent = append(vdfContent, ` "publishedfileid"	"`+item.WorkshopID+`"`)
	vdfContent = append(vdfContent, ` "contentfolder"	"`+absPath+`"`)
	vdfContent = append(vdfContent, ` "visibility"		"`+item.Visibility+`"`)
	vdfContent = append(vdfContent, ` "previewfile"		"`+absPath+`\thumb.png"`)

	// Fetch workshop info from the mod directory
	foundReadme, title, desc := fetch.Readme(item.ContentFolder)
	debug.Print("Locating Workshop Info at", item.ContentFolder, ":", foundReadme, title, desc)
	if foundReadme {
		item.Title = title
		item.Description = desc
	}

	// Append title and description if present
	if item.Title != "" {
		vdfContent = append(vdfContent, ` "title"		"`+item.Title+`"`)
	}
	if item.Description != "" {
		vdfContent = append(vdfContent, ` "description"		"`+item.Description+`"`)
	}

	// Fetch changelog from the mod directory
	foundChangelog, changelog := fetch.Changelog(item.ContentFolder)
	debug.Print("Locating Changelog at", item.ContentFolder, ":", foundChangelog, changelog)
	if foundChangelog {
		item.ChangeNote = changelog
	}

	// Append changelog if present
	if item.ChangeNote != "" {
		vdfContent = append(vdfContent, ` "changenote"	"`+item.ChangeNote+`"`)
	}

	// Append the closing brace
	vdfContent = append(vdfContent, `}`)

	// Join all parts to form the final VDF content
	return strings.Join(vdfContent, "\n")
}

// ReadVDF reads a VDF file from the given path and returns a VDFItem with the extracted fields.
func Read(vdfPath string) (VDFItem, error) {
	// Initialize an empty VDFItem
	item := VDFItem{}

	// Open the VDF file
	file, err := os.Open(vdfPath)
	if err != nil {
		return item, err
	}
	defer file.Close()

	// Create a scanner to read through the file line by line
	scanner := bufio.NewScanner(file)

	// Map to store extracted fields
	vdfFields := map[string]string{
		"appid":           "",
		"publishedfileid": "",
		"contentfolder":   "",
		"visibility":      "",
		"previewfile":     "",
		"title":           "",
		"description":     "",
		"changenote":      "",
	}

	// Read the VDF file line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Split the line into key and value, assuming VDF format: "key" "value"
		if strings.Contains(line, `"`) {
			parts := strings.SplitN(line, `"`, 3)
			if len(parts) >= 3 {
				key := strings.TrimSpace(parts[1])
				value := strings.TrimSpace(parts[2])
				value = strings.Trim(value, `"`)
				vdfFields[key] = value
			}
		}
	}

	// Map the extracted fields to the VDFItem struct
	item.AppID = vdfFields["appid"]
	item.WorkshopID = vdfFields["publishedfileid"]
	item.ContentFolder = vdfFields["contentfolder"]
	item.Visibility = vdfFields["visibility"]
	item.PreviewFile = vdfFields["previewfile"]
	item.Title = vdfFields["title"]
	item.Description = vdfFields["description"]
	item.ChangeNote = vdfFields["changenote"]

	// Return the populated VDFItem and any error from scanning
	if err := scanner.Err(); err != nil {
		return item, err
	}

	return item, nil
}
