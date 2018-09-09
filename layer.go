package tmx

import "encoding/xml"

// Layer is a layer of the map
type Layer struct {
	// Name is the name of the layer
	Name string `xml:"name,attr"`
	// X is the x coordinate of the layer
	X float64 `xml:"x,attr"`
	// Y is the y coordinate of the layer
	Y float64 `xml:"y,attr"`
	// Width is the width of the layer in tiles. Always the same as the map
	// width for fixed-size maps.
	Width int `xml:"width,attr"`
	// Height is the height of the layer in tiles. Always the same as the map
	// height for fixed-size maps.
	Height int `xml:"height,attr"`
	// Opacity is the opacity of the layer as a value from 0 to 1. Defaults to 1.
	Opacity float64 `xml:"opacity,attr"`
	// Visible is whether the layer is shown(1) or hidden(0). Defaults to 1.
	Visible int `xml:"visible,attr"`
	// OffsetX is the rendering offset for this layer in pixels.
	OffsetX float64 `xml:"offsetx,attr"`
	// OffsetY is the rendering offset for this layer in pixels.
	OffsetY float64 `xml:"offsety,attr"`
	// Properties are the properties of the layer
	Properties []Property `xml:"properties>property"`
	// Data is any data for the layer
	Data []Data `xml:"data"`
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (l *Layer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type layer Layer
	la := layer{
		Opacity: 1,
		Visible: 1,
	}
	if err := d.DecodeElement(&la, &start); err != nil {
		return err
	}
	*l = (Layer)(la)
	return nil
}
