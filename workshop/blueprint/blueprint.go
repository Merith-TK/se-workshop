package blueprint

import (
	"strings"

	"github.com/Merith-TK/se-workshop/fetch"
	"github.com/Merith-TK/utils/debug"
)

var (
	Files = []string{
		"bp.sbc",
	}
	VDFTemplate = []string{
		`"workshopitem"`,
		`{`,
		` "appid"		"244850"`,
		` "publishedfileid"	"{WORKSHOP_ID}"`,
		` "contentfolder"	"{MODPATH}"`,
		` "visibility"		"0"`,
		` "previewfile"		"{MODPATH}\thumb.png"`,
	}
)

func NewVDF(workshopID, modPath string) string {
	// Basic VDF template
	newVDF := VDFTemplate
	for i, line := range newVDF {
		newVDF[i] = strings.Replace(line, "{WORKSHOP_ID}", workshopID, -1)
		newVDF[i] = strings.Replace(newVDF[i], "{MODPATH}", modPath, -1)
	}

	// Fetch workshop info
	foundWork, title, desc := fetch.Readme(modPath)
	debug.Print("Locating Workshop Info at", modPath, ":", foundWork, title, desc)
	if foundWork {
		newVDF = append(newVDF, ` "title"		"`+title+`"`)
		newVDF = append(newVDF, ` "description"		"`+desc+`"`)
	}

	// Fetch changelog
	foundChangelog, changelog := fetch.Changelog(modPath)
	debug.Print("Locating Changelog at", modPath, ":", foundChangelog, changelog)
	if foundChangelog {
		newVDF = append(newVDF, ` "changenote"		"`+changelog+`"`)
	}

	// Append final lines
	newVDF = append(newVDF, `}`)
	return strings.Join(newVDF, "\n")
}
