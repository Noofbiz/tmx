package tmx

// Property is any custom data added to elements of the map
type Property struct {
	// Name is the name of the property
	Name string `xml:"name,attr"`
	// Type is the type of the property. It can be string, int, float, bool,
	// color or file
	Type string `xml:"type,attr"`
	// Value is the value of the property
	Value string `xml:"value,attr"`
}
