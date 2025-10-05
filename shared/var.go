package shared

import (
	"os"
	"path/filepath"
)

type ConstantsType struct {
	Appdata string
	Dir     struct {
		BP     string
		Mod    string
		Saves  string
		Script string
	}
}

var Constants ConstantsType

func init() {
	appdata := os.Getenv("APPDATA")
	Constants = ConstantsType{
		Appdata: filepath.Join(appdata, "SpaceEngineers") + string(filepath.Separator),
		Dir: struct {
			BP     string
			Mod    string
			Saves  string
			Script string
		}{
			BP:     filepath.Join(appdata, "SpaceEngineers", "Blueprints", "local") + string(filepath.Separator),
			Mod:    filepath.Join(appdata, "SpaceEngineers", "IngameScripts", "local") + string(filepath.Separator),
			Saves:  filepath.Join(appdata, "SpaceEngineers", "Saves") + string(filepath.Separator),
			Script: filepath.Join(appdata, "SpaceEngineers", "Scripts") + string(filepath.Separator),
		},
	}
}
