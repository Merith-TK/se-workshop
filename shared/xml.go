package shared

// WorkshopIds represents a collection of WorkshopId elements.
type WorkshopIds struct {
	WorkshopId WorkshopId `xml:"WorkshopId,omitempty"` // A single WorkshopId element.
}

// WorkshopId represents a single WorkshopId with an ID and service name.
type WorkshopId struct {
	Id          int64   `xml:"Id,omitempty"`          // The ID of the workshop item.
	ServiceName *string `xml:"ServiceName,omitempty"` // The name of the service (e.g., Steam).
}

// Header represents the XML header. It is used in other structs to manage XML metadata.
// This should NOT be used in the Metadata struct as it's not yet implemented for Entries using Header.
type Header struct {
	Text *string `xml:",chardata"`          // The text contained within the header.
	Nil  *string `xml:"nil,attr,omitempty"` // A flag indicating if the header is nil.
}

// Vector2 represents a 2D vector with X and Y coordinates.
type Vector2 struct {
	Header          // Embeds the Header struct for metadata.
	X      *float64 `xml:"x,attr,omitempty"` // X coordinate.
	Y      *float64 `xml:"y,attr,omitempty"` // Y coordinate.
}

// Vector3 represents a 3D vector, extending Vector2 with a Z coordinate.
type Vector3 struct {
	Vector2          // Embeds Vector2 for X and Y coordinates.
	Z       *float64 `xml:"z,attr,omitempty"` // Z coordinate.
}

// Vector4 represents a 4D vector, extending Vector3 with a W coordinate.
type Vector4 struct {
	Vector3          // Embeds Vector3 for X, Y, and Z coordinates.
	W       *float64 `xml:"w,attr,omitempty"` // W coordinate.
}

// PositionAndOrientation represents an object's position and orientation in 3D space.
type PositionAndOrientation struct {
	Header               // Embeds Header for XML metadata.
	Position    *Vector3 `xml:"Position,omitempty"`    // The position in 3D space.
	Forward     *Vector3 `xml:"Forward,omitempty"`     // The forward direction vector.
	Up          *Vector3 `xml:"Up,omitempty"`          // The up direction vector.
	Orientation *Vector4 `xml:"Orientation,omitempty"` // The orientation in 4D space (quaternion).
}

// Color represents a color in RGB or RGBA format, extending Vector3 for RGB channels.
type Color struct {
	Vector3             // Embeds Vector3 for RGB channels.
	PackedValue *string `xml:"PackedValue,omitempty"` // Packed value for the color (e.g., hexadecimal).
	R           *string `xml:"R,omitempty"`           // Red channel.
	G           *string `xml:"G,omitempty"`           // Green channel.
	B           *string `xml:"B,omitempty"`           // Blue channel.
	A           *string `xml:"A,omitempty"`           // Alpha (transparency) channel.
}

// ColorRGBA represents a color using separate RGBA values.
type ColorRGBA struct {
	Header         // Embeds Header for XML metadata.
	R      *string `xml:"R,attr,omitempty"` // Red channel.
	G      *string `xml:"G,attr,omitempty"` // Green channel.
	B      *string `xml:"B,attr,omitempty"` // Blue channel.
	A      *string `xml:"A,attr,omitempty"` // Alpha (transparency) channel.
}

// Slot represents an inventory slot, containing an item and metadata.
type Slot struct {
	Header         // Embeds Header for XML metadata.
	Index  *string `xml:"Index,omitempty"` // The index of the slot.
	Item   *string `xml:"Item,omitempty"`  // The item contained in the slot.
	Data   struct {
		Text          *string `xml:",chardata"`               // Text data associated with the slot.
		Type          *string `xml:"type,attr,omitempty"`     // Type of the slot (e.g., item type).
		Action        *string `xml:"Action,omitempty"`        // Action associated with the slot.
		BlockEntityId *string `xml:"BlockEntityId,omitempty"` // ID of the block entity, if applicable.
	} `xml:"Data,omitempty"` // Additional data associated with the slot.
}
