package semod

import (
	"encoding/xml"

	"github.com/Merith-TK/se-workshop/shared"
)

// Metadata represents the structure of modinfo.sbmi
type Metadata struct {
	XMLName      xml.Name                 `xml:"MyObjectBuilder_ModInfo"`
	SteamIDOwner string                   `xml:"SteamIDOwner"`
	WorkshopId   string                   `xml:"WorkshopId"`
	WorkshopIds  []shared.WorkshopIDEntry `xml:"WorkshopIds>WorkshopId"` // Correctly handles nested WorkshopId entries
}
