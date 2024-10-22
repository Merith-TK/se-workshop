package main

import (
	"os"
	"path/filepath"

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
			bpSbcPath := filepath.Join(fullpath, "bp.sbc")
			if _, err := os.Stat(bpSbcPath); err == nil {
				newvdf(fullpath)
			} else {
				continue
			}

			steamcmd("+workshop_build_item", filepath.Join(fullpath, "workshop.vdf"), "+quit")
			fixWorkshopID(fullpath)
		}

	}

	return nil
}

func newvdf(path string) (string, error) {
	bpSbcPath := filepath.Join(path, "bp.sbc")
	vdfString := ""
	if _, err := os.Stat(bpSbcPath); err == nil {
		vdfString = buildVDF(getWorkshopID(bpSbcPath), path)
		workshopVdfPath := filepath.Join(path, "workshop.vdf")
		file, err := os.OpenFile(workshopVdfPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return vdfString, err
		}
		defer file.Close()
		_, err = file.WriteString(vdfString)
		if err != nil {
			return "", err
		}
	}
	return vdfString, nil
}
