package main

import (
	"flag"
	"fmt"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/se-workshop/utils/sebp"
	"github.com/Merith-TK/utils/debug"
	"github.com/beevik/etree"
	"github.com/dlclark/regexp2"
)

var confFile = flag.String("conf", "chopshop.json", "configuration file")
var Config GridConfig

func main() {
	flag.Parse()
	if *confFile != "" || !shared.FileExists(*confFile) {
		var err error
		Config, err = ReadConf(*confFile)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(Config)
	for _, v := range flag.Args() {
		found, path := sebp.LocateBP(v)
		if found {
			chopshop(path)
		} else {
			fmt.Println("Not a valid blueprint path:", v)
		}
	}
}

func chopshop(path string) {
	if !shared.FileExists(path) {
		fmt.Println("File not found:", path)
		return
	}

	fmt.Println("Processing file:", path)
	// Load the XML file to map `shared.BPData` to `etree.Document`
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var loadedMap []GridMapping
	for _, gridSize := range doc.FindElements("Definitions/ShipBlueprints/ShipBlueprint/CubeGrids/CubeGrid/GridSizeEnum") {
		if gridSize.Text() == "Small" {
			loadedMap = Config.SmallGrid
			fmt.Println("Parsing Small Grid")
			break
		}
		if gridSize.Text() == "Large" {
			loadedMap = Config.LargeGrid
			fmt.Println("Parsing Large Grid")
			break
		}
	}

	Blocks := doc.FindElements("//CubeBlocks/MyObjectBuilder_CubeBlock")
	Altered := 0
	// Process the document to replace blocks
	for _, block := range Blocks {
		debug.SetTitle("BLOCK")
		targetBlock := block.SelectElement("SubtypeName")
		debug.Print(targetBlock.Text())
		for _, mapping := range loadedMap {
			debug.SetTitle("MAP  ")
			debug.Print(mapping.Repl)
			re := regexp2.MustCompile(mapping.Repl, regexp2.RE2)
			if match, _ := re.MatchString(targetBlock.Text()); match {
				debug.SetTitle("MATCH")
				newText, _ := re.Replace(targetBlock.Text(), mapping.With, -1, -1)
				targetBlock.SetText(newText)
				Altered++
				debug.Print(newText)
			} else {
				debug.SetTitle("NOMAT")
				debug.Print(targetBlock.Text())
			}
		}
	}
	fmt.Printf("Total blocks replaced: %d\n", Altered)

	// Save the modified document back to the file
	doc.Indent(2)
	// if err := doc.WriteToFile(path + ".xml"); err != nil {
	if err := doc.WriteToFile(path); err != nil {
		fmt.Println("Error writing file:", err)
	}
	fmt.Println("File processed successfully:", path)
}
