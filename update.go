package main

import (
	"os"
	"path/filepath"

	"github.com/Merith-TK/se-workshop/steam"
	"github.com/Merith-TK/se-workshop/workshop/blueprint"

	"github.com/Merith-TK/utils/debug"
)

func update(args ...string) error {
	debug.Print("Updating", args)
	for _, arg := range args {
		debug.Print("Updating", arg)
		fullpath, abserr := filepath.Abs(arg)
		if abserr != nil {
			debug.Print("Failed to get absolute path:", abserr)
			return abserr
		}

		// Check if the path is a directory
		fileinfo, staterr := os.Stat(fullpath)
		if staterr != nil {
			debug.Print("Failed to get file info:", staterr)
			return staterr
		}

		if fileinfo.IsDir() {
			steam.CMD("+workshop_build_item", filepath.Join(fullpath, "workshop.vdf"), "+quit")
			bpSbcPath := filepath.Join(fullpath, "bp.sbc")
			if _, err := os.Stat(bpSbcPath); err == nil {
				blueprint.WorkshopID(fullpath)
			}

		}

	}

	return nil
}
