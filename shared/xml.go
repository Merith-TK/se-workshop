package shared

// WorkshopIds represents the WorkshopIds element containing WorkshopId elements.
type WorkshopIds struct {
	WorkshopId WorkshopId `xml:"WorkshopId,omitempty"`
}

// WorkshopId represents a single WorkshopId element.
type WorkshopId struct {
	Id          int64   `xml:"Id,omitempty"`
	ServiceName *string `xml:"ServiceName,omitempty"`
}

// Usage of Header inside the Metadata struct
// should NOT be used, Entries that use Header in the Metadata
// struct are not yet implemented.
type Header struct {
	Text *string `xml:",chardata"`
	Nil  *string `xml:"nil,attr,omitempty"`
}
type Vector2 struct {
	Header
	X *float64 `xml:"x,attr,omitempty"`
	Y *float64 `xml:"y,attr,omitempty"`
}

type Vector3 struct {
	Vector2
	Z *float64 `xml:"z,attr,omitempty"`
}

type Vector4 struct {
	Vector3
	W *float64 `xml:"w,attr,omitempty"`
}

type PositionAndOrientation struct {
	Header
	Position    *Vector3 `xml:"Position,omitempty"`
	Forward     *Vector3 `xml:"Forward,omitempty"`
	Up          *Vector3 `xml:"Up,omitempty"`
	Orientation *Vector4 `xml:"Orientation,omitempty"`
}

type Color struct {
	Vector3
	PackedValue *string `xml:"PackedValue,omitempty"`
	R           *string `xml:"R,omitempty"`
	G           *string `xml:"G,omitempty"`
	B           *string `xml:"B,omitempty"`
	A           *string `xml:"A,omitempty"`
}

type ColorRGBA struct {
	Header
	R *string `xml:"R,attr,omitempty"`
	G *string `xml:"G,attr,omitempty"`
	B *string `xml:"B,attr,omitempty"`
	A *string `xml:"A,attr,omitempty"`
}

type Slot struct {
	Header
	Index *string `xml:"Index,omitempty"`
	Item  *string `xml:"Item,omitempty"`
	Data  struct {
		Text          *string `xml:",chardata"`
		Type          *string `xml:"type,attr,omitempty"`
		Action        *string `xml:"Action,omitempty"`
		BlockEntityId *string `xml:"BlockEntityId,omitempty"`
	} `xml:"Data,omitempty"`
}
