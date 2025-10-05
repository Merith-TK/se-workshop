package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/debug"
	"github.com/beevik/etree"
)

// FileExists checks whether a file exists at the specified path.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// PrintHelp prints the help message for the CLI commands.
// If msg is not empty, it will be printed before the usage instructions.
func PrintHelp(msg string) {
	if msg != "" {
		println("Error:", msg)
		println()
	}
	
	println("sew - Space Engineers Workshop Tool")
	println("A command-line tool for managing Steam Workshop items for Space Engineers")
	println()
	println("USAGE:")
	println("  sew <command> [args...]")
	println()
	println("COMMANDS:")
	println()
	println("  Authentication:")
	println("    login <username> [password] [steamauth]")
	println("        Log into SteamCMD (required before uploading)")
	println("        Password and steam auth are optional - will prompt if needed")
	println()
	println("  Workshop Management:")
	println("    upload [path]")
	println("        Upload or update a workshop item (uses current dir if no path)")
	println("    download <workshop_id>")
	println("        Download a workshop item (e.g: sew download 123456789)")
	println("    get-id [path]")
	println("        Get the workshop URL for an item (uses current dir if no path)")
	println("    set-id <path> <workshop_id>")
	println("        Set the workshop ID for a blueprint or mod")
	println()
	println("  Blueprint Commands:")
	println("    bp folder")
	println("        Show the path to your local blueprints folder")
	println()
	println("  Utility Commands:")
	println("    cmd <args...>")
	println("        Run SteamCMD directly with the given arguments")
	println("    vent-steam")
	println("        Restart Steam (fixes offline status after using SteamCMD)")
	println("    help, ?")
	println("        Show this help message")
	println()
	println("EXAMPLES:")
	println("  sew login myusername")
	println("  sew upload C:\\path\\to\\blueprint")
	println("  sew download 123456789")
	println("  sew get-id .")
	println("  sew set-id . 123456789")
	println("  sew bp folder")
	println()
	println("For more information, see: https://github.com/Merith-TK/se-workshop")
}

// PWD returns the current working directory.
// If there is an error, it logs the error and returns an empty string.
func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		debug.Print("Error getting current directory:", err)
		return ""
	}
	return pwd
}

// GetWorkshopID retrieves the workshop ID from the specified path.
// The path can be a VDF file, modinfo file (.sbmi), or a blueprint file (.sbc).
// Returns "0" if no valid workshop ID is found.
func GetWorkshopID(path string) string {
	debug.SetTitle("Get Workshop ID")
	defer debug.ResetTitle()

	// Handle different file types and adjust path accordingly
	if !strings.HasSuffix(path, ".sbc") && !strings.HasSuffix(path, ".sbmi") {
		switch {
		case FileExists(filepath.Join(path, WorkshopVDFFileName)):
			path = filepath.Join(path, WorkshopVDFFileName)
		case FileExists(filepath.Join(path, ModInfoFileName)):
			path = filepath.Join(path, ModInfoFileName)
		case FileExists(filepath.Join(path, BlueprintFileName)):
			path = filepath.Join(path, BlueprintFileName)
		default:
			debug.Print("Invalid path:", path)
			return "0"
		}
	}

	// Ensure the file exists
	if !FileExists(path) {
		debug.Print("File not found:", path)
		return "0"
	}

	// Handle VDF files
	if strings.HasSuffix(path, WorkshopVDFFileName) {
		debug.Print("Reading VDF file:", path)
		vdfContent, err := vdf.Read(path)
		if err != nil {
			debug.Print("Error reading VDF file:", err)
			return "0"
		}
		return vdfContent.WorkshopID
	}

	// Handle modinfo files (.sbmi)
	if strings.HasSuffix(path, ".sbmi") {
		debug.Print("Reading modinfo file:", path)
		return parseModInfoWorkshopID(path)
	}

	// Handle blueprint files (.sbc)
	if strings.HasSuffix(path, ".sbc") {
		debug.Print("Reading blueprint file:", path)
		return parseBlueprintWorkshopID(path)
	}

	// Fallback to checking for a VDF file in the current directory
	if FileExists(WorkshopVDFFileName) {
		vdfContent, err := vdf.Read(WorkshopVDFFileName)
		if err != nil {
			debug.Print("Error reading VDF file:", err)
			return "0"
		}
		return vdfContent.WorkshopID
	}

	return "0"
}

// SetWorkshopID sets the workshop ID in the specified file.
// The path can be a modinfo file (.sbmi) or a blueprint file (.sbc).
func SetWorkshopID(path string, workshopID string) {
	if !strings.HasSuffix(path, ".sbc") && !strings.HasSuffix(path, ".sbmi") {
		if FileExists(filepath.Join(path, ModInfoFileName)) {
			path = filepath.Join(path, ModInfoFileName)
		} else if FileExists(filepath.Join(path, BlueprintFileName)) {
			path = filepath.Join(path, BlueprintFileName)
		} else {
			debug.Print("Invalid path:", path)
			return
		}
	}

	if !FileExists(path) {
		debug.Print("File not found:", path)
		return
	}

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		debug.Print("Error reading file:", err)
		return
	}

	if strings.HasSuffix(path, ".sbmi") {
		setModInfoWorkshopID(doc, workshopID)
	}

	if strings.HasSuffix(path, ".sbc") {
		setBlueprintWorkshopID(doc, workshopID)
	}

	// Write the updated XML document back to the file
	doc.Indent(2)
	if err := doc.WriteToFile(path); err != nil {
		debug.Print("Error writing file:", err)
	}
}

// CleanXML removes empty elements from the XML content and normalizes line endings.
func CleanXML(content string) string {
	// Trim trailing whitespace and normalize line endings
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t\r ")
	}
	updatedXML := strings.Join(lines, "\n")
	updatedXML = strings.ReplaceAll(updatedXML, "&#xA;", "")
	updatedXML = strings.ReplaceAll(updatedXML, "\r\n", "\n")

	// Parse and clean the XML document
	doc := etree.NewDocument()
	if err := doc.ReadFromString(updatedXML); err != nil {
		panic(err)
	}
	cleanEmptyElements(doc.Root())
	doc.Indent(2)
	cleanedXML, _ := doc.WriteToString()
	return cleanedXML
}

// cleanEmptyElements recursively removes empty elements from the XML tree.
func cleanEmptyElements(element *etree.Element) {
	for _, child := range element.ChildElements() {
		cleanEmptyElements(child)
	}
	if len(element.ChildElements()) == 0 && strings.TrimSpace(element.Text()) == "" && len(element.Attr) == 0 {
		element.Parent().RemoveChild(element)
	}
}

// Helper function to parse workshop ID from modinfo.sbmi files.
func parseModInfoWorkshopID(path string) string {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		debug.Print("Error reading file:", err)
		return "0"
	}
	root := doc.SelectElement("MyObjectBuilder_ModInfo")
	if root == nil {
		debug.Print("Invalid XML structure")
		return "0"
	}

	workshopIDElement := root.SelectElement("WorkshopId")
	if workshopIDElement == nil || workshopIDElement.Text() == "0" {
		workshopIds := doc.FindElements("//MyObjectBuilder_ModInfo/WorkshopIds/WorkshopId")
		for _, id := range workshopIds {
			idElement := id.SelectElement("Id")
			serviceNameElement := id.SelectElement("ServiceName")
			if idElement != nil && serviceNameElement != nil {
				fmt.Printf("WorkshopId: %s, ServiceName: %s\n", idElement.Text(), serviceNameElement.Text())
				return idElement.Text()
			}
		}
	}
	return workshopIDElement.Text()
}

// Helper function to parse workshop ID from blueprint files (.sbc).
func parseBlueprintWorkshopID(path string) string {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		debug.Print("Error reading file:", err)
		return "0"
	}
	root := doc.SelectElement("Definitions")
	if root == nil {
		debug.Print("Invalid XML structure")
		return "0"
	}
	shipBlueprint := root.SelectElement("ShipBlueprints").SelectElement("ShipBlueprint")
	if shipBlueprint == nil {
		debug.Print("ShipBlueprint element not found")
		return "0"
	}

	workshopIDElement := shipBlueprint.SelectElement("WorkshopId")
	if workshopIDElement == nil || workshopIDElement.Text() == "0" {
		workshopIds := doc.FindElements("//Definitions/ShipBlueprints/ShipBlueprint/WorkshopIds/WorkshopId")
		for _, id := range workshopIds {
			idElement := id.SelectElement("Id")
			serviceNameElement := id.SelectElement("ServiceName")
			if idElement != nil && serviceNameElement != nil {
				fmt.Printf("WorkshopId: %s, ServiceName: %s\n", idElement.Text(), serviceNameElement.Text())
				return idElement.Text()
			}
		}
	}
	return workshopIDElement.Text()
}

// Helper function to set workshop ID in modinfo.sbmi files.
func setModInfoWorkshopID(doc *etree.Document, workshopID string) {
	root := doc.SelectElement("MyObjectBuilder_ModInfo")
	if root == nil {
		debug.Print("Invalid XML structure")
		return
	}
	workshopIDElement := root.SelectElement("WorkshopId")
	if workshopIDElement == nil {
		debug.Print("WorkshopId element not found")
		return
	}
	workshopIDElement.SetText(workshopID)
}

// Helper function to set workshop ID in blueprint files (.sbc).
func setBlueprintWorkshopID(doc *etree.Document, workshopID string) {
	root := doc.SelectElement("Definitions")
	if root == nil {
		debug.Print("Invalid XML structure")
		return
	}
	shipBlueprint := root.SelectElement("ShipBlueprints").SelectElement("ShipBlueprint")
	if shipBlueprint == nil {
		debug.Print("ShipBlueprint element not found")
		return
	}
	workshopIDElement := shipBlueprint.SelectElement("WorkshopId")
	if workshopIDElement == nil {
		workshopIDElement = shipBlueprint.CreateElement("WorkshopId")
	}
	workshopIDElement.SetText(workshopID)
}
