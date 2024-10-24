// Package vdf provides functions to manage and manipulate VDF (Valve Data Format) files
// for Steam Workshop items. This includes building VDF files for uploading or updating
// workshop items, as well as reading existing VDF files.
package vdf

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/utils/debug"
)

// VDFItem represents the necessary fields required to create or read a VDF file.
// It includes the App ID, Workshop ID, content folder path, visibility, preview file path,
// title, description, and change note for the item.
type VDFItem struct {
	AppID         string // The Steam app ID (default: 244850 for Space Engineers)
	WorkshopID    string // The Steam Workshop ID of the item
	ContentFolder string // The absolute path to the content folder
	Visibility    string // Visibility of the item (0 = public, 1 = friends only, 2 = private)
	PreviewFile   string // The path to the preview image file
	Title         string // Title of the workshop item
	Description   string // Description of the workshop item
	ChangeNote    string // Changelog or update notes
}

// Build generates the content of a VDF file based on the provided VDFItem struct.
// This content is used to upload or update an item in the Steam Workshop.
//
// The function will automatically fill missing fields such as app ID (default 244850),
// visibility (default 0), and preview file (default: "thumb.png"). It also attempts to
// retrieve the title and description from an "info.txt" file in the content folder,
// and the changelog from a "changelog.txt" file if they exist.
//
// Parameters:
// - item: A VDFItem struct containing the necessary fields to build the VDF file.
//
// Returns:
// - A string containing the generated VDF file content, or an empty string in case of errors.
func Build(item VDFItem) string {
	absPath, err := filepath.Abs(item.ContentFolder)
	if err != nil {
		return ""
	}

	// Set default values if not provided
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
	vdfContent = append(vdfContent, ` "appid"		"`+item.AppID+`"`)
	vdfContent = append(vdfContent, ` "publishedfileid"	"`+item.WorkshopID+`"`)
	vdfContent = append(vdfContent, ` "contentfolder"	"`+absPath+`"`)
	vdfContent = append(vdfContent, ` "visibility"		"`+item.Visibility+`"`)
	vdfContent = append(vdfContent, ` "previewfile"		"`+item.PreviewFile+`"`)

	// Fetch workshop info (title and description) from "info.txt"
	foundReadme, title, desc := Readme(item.ContentFolder)
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

	// Fetch changelog from "changelog.txt"
	foundChangelog, changelog := Changelog(item.ContentFolder)
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

// Read reads a VDF file from the provided path and extracts the fields into a VDFItem struct.
//
// The VDF file is expected to contain fields such as appid, publishedfileid, contentfolder, etc.
// This function will scan through the file line by line, extract the values, and populate
// the VDFItem struct.
//
// Parameters:
// - vdfPath: The file path to the VDF file to read.
//
// Returns:
// - VDFItem: A populated VDFItem struct containing the extracted values.
// - error: An error if any issues occur while reading the file.
func Read(vdfPath string) (VDFItem, error) {
	// Initialize an empty VDFItem
	item := VDFItem{}

	// Open the VDF file
	file, err := os.Open(vdfPath)
	if err != nil {
		return item, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
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

	// Read the VDF file line by line and extract key-value pairs
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
