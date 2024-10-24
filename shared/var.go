package shared

import "os"

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
	Constants = ConstantsType{
		Appdata: os.Getenv("APPDATA") + "\\SpaceEngineers\\",
		Dir: struct {
			BP     string
			Mod    string
			Saves  string
			Script string
		}{
			BP:     os.Getenv("APPDATA") + "\\SpaceEngineers\\Blueprints\\local\\",
			Mod:    os.Getenv("APPDATA") + "\\SpaceEngineers\\IngameScripts\\local\\",
			Saves:  os.Getenv("APPDATA") + "\\SpaceEngineers\\Saves\\",
			Script: os.Getenv("APPDATA") + "\\SpaceEngineers\\Scripts\\",
		},
	}
}
