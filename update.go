package main

import (
	"os"
	"path/filepath"
)

func update(args ...string) error {
	for _, arg := range args {
		if _, err := os.Stat(filepath.Join(arg)); os.IsNotExist(err) {
			return err
		} else {
			readvdf(arg)
		}
	}

	return nil
}

// TODO: PATH_TO_STEAM_CMD\steamcmd.exe +login YOUR_ACCOUNT_NAME +workshop_build_item "FULL_PATH_TO\UpdateContainer.vdf" +logoff +quit

func readvdf(path string) error {
	bpSbcPath := filepath.Join(path, "bp.sbc")
	if _, err := os.Stat(bpSbcPath); err == nil {
		workshopVdfPath := filepath.Join(path, "workshop.vdf")
		if _, err := os.Stat(workshopVdfPath); os.IsNotExist(err) {
			file, err := os.Create(workshopVdfPath)
			if err != nil {
				return err
			}
			defer file.Close()
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			_, err = file.WriteString(buildVDF(getWorkshopID(bpSbcPath), absPath))
		}
	}
	return nil
}
