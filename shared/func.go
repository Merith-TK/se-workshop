package shared

import (
	"fmt"
	"os"
	"strings"

	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/debug"
	"github.com/beevik/etree"
)

// fileExists checks whether a file exists and is accessible
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PrintHelp(msg string) {

	if msg != "" {
		println(msg)
	}

	println("Usage: se-workshop <command> [args...]")
	println("Commands:")
	println("  build-vdf <path> - Build a VDF file for the provided blueprint path")
	println("  folder - Print the current blueprint directory")
	println("  get-id <path> - Get the workshop ID for the provided blueprint path")

	println("This is a incomplete help message.")
}

func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		debug.Print("Error getting current directory:", err)
		return ""
	}
	return pwd
}

func GetWorkshopID(path string) string {
	debug.SetTitle("Get Workshop ID")
	defer debug.ResetTitle()

	if !strings.HasSuffix(path, ".sbc") && !strings.HasSuffix(path, ".sbmi") {
		if FileExists(path + "\\workshop.vdf") {
			path = path + "\\workshop.vdf"
		} else if FileExists(path + "\\modinfo.sbmi") {
			path = path + "\\modinfo.sbmi"
		} else if FileExists(path + "\\bp.sbc") {
			path = path + "\\bp.sbc"
		} else {
			debug.Print("Invalid path:", path)
			return "0"
		}
	}

	if !FileExists(path) {
		debug.Print("File not found:", path)
		return "0"
	}

	if strings.HasSuffix(path, "workshop.vdf") {
		debug.Print("Reading VDF file:", path)
		vdfContent, err := vdf.Read(path)
		if err != nil {
			debug.Print("Error reading VDF file:", err)
			return "0"
		}
		return vdfContent.WorkshopID
	}
	if strings.HasSuffix(path, ".sbmi") {
		debug.Print("Reading modinfo file:", path)
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
		if workshopIDElement == nil {
			debug.Print("WorkshopId element not found")
			return "0"
		}

		if workshopIDElement.Text() == "0" {
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

	if strings.HasSuffix(path, ".sbc") {
		debug.Print("Reading blueprint file:", path)
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

	if FileExists("workshop.vdf") {
		vdfContent, err := vdf.Read("workshop.vdf")
		if err != nil {
			debug.Print("Error reading VDF file:", err)
			return "0"
		}
		return vdfContent.WorkshopID
	}

	return "0"
}

func SetWorkshopID(path string, workshopID string) {
	if !strings.HasSuffix(path, ".sbc") && !strings.HasSuffix(path, ".sbmi") {
		if FileExists(path + "\\modinfo.sbmi") {
			path = path + "\\modinfo.sbmi"
		} else if FileExists(path + "\\bp.sbc") {
			path = path + "\\bp.sbc"
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

		if workshopIDElement.Text() == "0" {
			workshopIds := doc.FindElements("//MyObjectBuilder_ModInfo/WorkshopIds/WorkshopId")
			for _, id := range workshopIds {
				idElement := id.SelectElement("Id")
				serviceNameElement := id.SelectElement("ServiceName")
				if idElement != nil && serviceNameElement != nil {
					fmt.Printf("WorkshopId: %s, ServiceName: %s\n", idElement.Text(), serviceNameElement.Text())
					if serviceNameElement.Text() == "Steam" {
						idElement.SetText(workshopID)
					}
				}
			}
		}

		workshopIDElement.SetText(workshopID)
	}

	if strings.HasSuffix(path, ".sbc") {
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
			debug.Print("Blueprint WorkshopId element not found")
			workshopIds := doc.FindElements("//Definitions/ShipBlueprints/ShipBlueprint/WorkshopIds/WorkshopId")
			var serviceNameElement *etree.Element
			for _, id := range workshopIds {
				idElement := id.SelectElement("Id")
				serviceNameElement = id.SelectElement("ServiceName")
				if idElement != nil && serviceNameElement != nil {
					if serviceNameElement.Text() == "Steam" {
						idElement.SetText(workshopID)
					}
				}
			}

			if serviceNameElement == nil {
				shipBlueprint.CreateElement("WorkshopIds").CreateElement("WorkshopId").CreateElement("Id").SetText(workshopID)
				shipBlueprint.FindElement("//Definitions/ShipBlueprints/ShipBlueprint/WorkshopIds/WorkshopId").CreateElement("ServiceName").SetText("Steam")
			}
		} else {
			debug.Print("Setting Workshop ID:", workshopID)
			workshopIDElement.SetText(workshopID)
		}
	}

	doc.Indent(2)
	if err := doc.WriteToFile(path); err != nil {
		debug.Print("Error writing file:", err)
	}
}

// CleanXML removes empty elements from the XML content and normalizes line endings
func CleanXML(content string) string {
	// Split the content into lines
	lines := strings.Split(content, "\n")

	// Trim the trailing whitespace from each line
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
		lines[i] = strings.TrimRight(line, "\r")
		lines[i] = strings.TrimRight(line, " ")
	}

	// Join the lines back together with newline characters
	updatedXML := strings.Join(lines, "\n")
	// Remove extra newline and carriage return characters
	updatedXML = strings.ReplaceAll(updatedXML, "&#xA;", "")
	updatedXML = strings.ReplaceAll(updatedXML, "\r\n", "\n") // normalize line endings

	doc := etree.NewDocument()
	if err := doc.ReadFromString(updatedXML); err != nil {
		panic(err)
	}

	cleanEmptyElements(doc.Root())
	doc.Indent(2)
	cleanedXML, _ := doc.WriteToString()
	return cleanedXML
}

// Recursively clean empty elements from the XML tree
func cleanEmptyElements(element *etree.Element) {
	// Iterate through all child elements
	for _, child := range element.ChildElements() {
		cleanEmptyElements(child) // Recursively clean child elements
	}

	// Check if the current element is empty (no child elements, no text, and no attributes)
	if len(element.ChildElements()) == 0 && strings.TrimSpace(element.Text()) == "" && len(element.Attr) == 0 {
		element.Parent().RemoveChild(element)
	}
}
