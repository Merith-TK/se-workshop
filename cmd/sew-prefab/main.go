package main

import (
	"flag"
	"fmt"

	"github.com/beevik/etree"
)

var (
	toBlueprint = flag.Bool("to-bp", false, "Convert prefab to blueprint (default behavior)")
	toMod       = flag.Bool("to-mod", false, "Convert blueprint to spawn mod (not yet implemented)")
	outputDir   = flag.String("output", "", "Output directory for conversion (required for -to-mod)")
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		printUsage()
		return
	}

	// Validate flags
	if *toBlueprint && *toMod {
		fmt.Println("Error: Cannot specify both -to-bp and -to-mod flags")
		return
	}

	if *toMod && *outputDir == "" {
		fmt.Println("Error: -output directory is required when using -to-mod")
		return
	}

	// Default behavior is to convert to blueprint if no flags specified
	if !*toMod {
		*toBlueprint = true
	}

	if *toBlueprint {
		convertToBlueprint(flag.Args())
	} else if *toMod {
		convertToMod(flag.Args())
	}
}

func printUsage() {
	fmt.Println("sew-prefab - Convert between Space Engineers blueprints and prefabs")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  sew-prefab [flags] <file1> [file2] ...")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -to-bp              Convert prefab to blueprint (default)")
	fmt.Println("  -to-mod             Convert blueprint to spawn mod (not yet implemented)")
	fmt.Println("  -output <dir>       Output directory (for future mod conversion)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  sew-prefab prefab.sbc                    # Convert prefab to blueprint")
	fmt.Println("  sew-prefab -to-bp prefab.sbc             # Same as above")
	fmt.Println("  sew-prefab -to-mod -output ./mod bp.sbc  # Blueprint to mod (coming soon)")
}

func convertToBlueprint(filePaths []string) {
	for _, filePath := range filePaths {
		if err := prefabToBlueprint(filePath); err != nil {
			fmt.Printf("Error converting %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Successfully converted %s to blueprint\n", filePath)
		}
	}
}

func convertToMod(filePaths []string) {
	if len(filePaths) != 1 {
		fmt.Println("Error: -to-mod mode accepts exactly one blueprint file")
		return
	}

	fmt.Println("Blueprint to mod conversion is not yet implemented.")
	fmt.Println("Please use the current tool for prefab to blueprint conversion only.")
}

func prefabToBlueprint(filePath string) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(filePath); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	prefab := doc.FindElement("//Prefab/CubeGrids")
	if prefab == nil {
		return fmt.Errorf("no CubeGrids element found in prefab")
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
	return newDoc.WriteToFile(filePath)
}
