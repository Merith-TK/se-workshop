package semod

import (
	"encoding/xml"

	"github.com/Merith-TK/se-workshop/shared"
)

type Metadata struct {
	XMLName      xml.Name                `xml:"MyObjectBuilder_ModInfo"`
	Text         string                  `xml:",chardata"`
	Xsd          string                  `xml:"xsd,attr,omitempty"`
	Xsi          string                  `xml:"xsi,attr,omitempty"`
	SteamIDOwner string                  `xml:"SteamIDOwner,omitempty"`
	WorkshopId   string                  `xml:"WorkshopId,omitempty"`
	WorkshopIds  []shared.WorkshopIDItem `xml:"WorkshopIds,omitempty"`
}
