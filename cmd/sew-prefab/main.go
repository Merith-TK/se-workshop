package main

import (
	"flag"
	"fmt"

	"github.com/beevik/etree"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Please provide at least one file as an argument.")
		return
	}

	for _, filePath := range flag.Args() {
		doc := etree.NewDocument()
		if err := doc.ReadFromFile(filePath); err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		prefab := doc.FindElement("//Prefab/CubeGrids")
		if prefab == nil {
			fmt.Printf("No CubeGrids element found in file %s\n", filePath)
			continue
		}

		newDoc := etree.NewDocument()
		newDoc.CreateProcInst("xml", `version="1.0"`)
		definitions := newDoc.CreateElement("Definitions")
		definitions.CreateAttr("xmlns:xsd", "http://www.w3.org/2001/XMLSchema")
		definitions.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
		shipBlueprints := definitions.CreateElement("ShipBlueprints")
		shipBlueprint := shipBlueprints.CreateElement("ShipBlueprint")
		shipBlueprint.CreateAttr("xsi:type", "MyObjectBuilder_ShipBlueprintDefinition")
		shipBlueprint.AddChild(prefab)

		newDoc.Indent(2)
		if err := newDoc.WriteToFile(filePath); err != nil {
			fmt.Printf("Error writing file %s: %v\n", filePath, err)
			continue
		}

		fmt.Printf("Successfully transformed file %s\n", filePath)
	}
}
