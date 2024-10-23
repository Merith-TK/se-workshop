package sebp

import (
	"os"
	"path/filepath"

	"github.com/Merith-TK/se-workshop/steam"
	"github.com/Merith-TK/se-workshop/vdf"
	"github.com/Merith-TK/utils/debug"
)

func Upload(args ...string) error {
	debug.SetTitle("Upload")
	defer debug.ResetTitle()

	for _, arg := range args {
		debug.Print("Uploading", arg)
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
			vdfitem := vdf.VDFItem{
				ContentFolder: fullpath,
				WorkshopID:    WorkshopID(fullpath),
			}

			vdfpath := filepath.Join(fullpath, "workshop.vdf")
			err := os.WriteFile(vdfpath, []byte(vdf.Build(vdfitem)), 0644)
			if err != nil {
				debug.Print("Failed to write workshop.vdf:", err)
				return err
			}
			steam.CMD("+workshop_build_item", vdfpath, "+quit")

		}

	}
	return nil
}
