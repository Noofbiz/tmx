package tmx

import "encoding/xml"

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

func (p *Property) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	prop := struct {
		// Name is the name of the property
		Name string `xml:"name,attr"`
		// Type is the type of the property. It can be string, int, float, bool,
		// color or file
		Type string `xml:"type,attr"`
		// Value is the value of the property
		Value string `xml:"value,attr"`

		CharData string `xml:",chardata"`
	}{}
	if err := d.DecodeElement(&prop, &start); err != nil {
		return err
	}

	p.Name = prop.Name
	p.Type = prop.Type
	p.Value = prop.Value
	if len(p.Value) == 0 {
		p.Value = prop.CharData
	}

	return nil

}
