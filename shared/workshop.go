package shared

type WorkshopIDItem struct {
	Text       string `xml:",chardata"`
	WorkshopId struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"Id,omitempty"`
		ServiceName string `xml:"ServiceName,omitempty"`
	} `xml:"WorkshopId,omitempty"`
}
