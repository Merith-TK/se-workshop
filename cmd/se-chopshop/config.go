package main

import (
	"fmt"
	"os"

	"github.com/Merith-TK/se-workshop/shared"
	"github.com/Merith-TK/utils/debug"
	"github.com/dlclark/regexp2"
	json "github.com/yosuke-furukawa/json5/encoding/json5"
)

type GridConfig struct {
	SmallGrid []GridMapping `json:"smallgrid"`
	LargeGrid []GridMapping `json:"largegrid"`
}

type GridMapping struct {
	Repl string `json:"repl"`
	With string `json:"with"`
}

// ReadConf reads and parses the configuration file
func ReadConf(file string) (GridConfig, error) {
	if !shared.FileExists(file) {
		return GridConfig{}, fmt.Errorf("config file not found: %s", file)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return GridConfig{}, fmt.Errorf("failed to read file %s: %w", file, err)
	}

	var config GridConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return GridConfig{}, fmt.Errorf("failed to parse JSON from %s: %w", file, err)
	}

	jsonString, _ := json.MarshalIndent(config, "", "  ")
	debug.Print("Config:\n", string(jsonString), "\n")

	for _, mapping := range append(config.SmallGrid, config.LargeGrid...) {
		if _, err := regexp2.Compile(mapping.Repl, regexp2.RE2); err != nil {
			return GridConfig{}, fmt.Errorf("invalid regex in config: %s, error: %w", mapping.Repl, err)
		}
	}
	return config, nil
}
