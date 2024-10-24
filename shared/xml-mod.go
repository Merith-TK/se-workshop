package shared

import (
	"encoding/xml"
)

// Metadata represents the structure of modinfo.sbmi
type MetadataMod struct {
	XMLName      xml.Name     `xml:"MyObjectBuilder_ModInfo"`
	SteamIDOwner string       `xml:"SteamIDOwner"`
	WorkshopId   string       `xml:"WorkshopId"`
	WorkshopIds  []WorkshopId `xml:"WorkshopIds>WorkshopId"` // Correctly handles nested WorkshopId entries
}
