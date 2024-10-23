package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Merith-TK/se-workshop/workshop/sebp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the XML file path as an argument")
		return
	}

	filePath := os.Args[1]
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var metadata sebp.Metadata
	err = xml.Unmarshal(byteValue, &metadata)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	fmt.Printf("Parsed Metadata: %+v", metadata)
}
