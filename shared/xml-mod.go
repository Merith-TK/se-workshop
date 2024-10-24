package shared

import (
	"encoding/xml"
)

// MetadataMod represents the structure of the modinfo.sbmi file.
// It contains details about a mod, including the owner and associated Workshop IDs.
type MetadataMod struct {
	XMLName      xml.Name     `xml:"MyObjectBuilder_ModInfo"` // The XML name of the root element.
	SteamIDOwner string       `xml:"SteamIDOwner"`            // The Steam ID of the mod's owner.
	WorkshopId   string       `xml:"WorkshopId"`              // The primary Workshop ID for the mod.
	WorkshopIds  []WorkshopId `xml:"WorkshopIds>WorkshopId"`  // A list of additional Workshop IDs, handling nested WorkshopId entries correctly.
}
