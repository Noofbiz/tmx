package tmx

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
)

// ObjectGroup is a group of objects
type ObjectGroup struct {
	// Name is the name of the object group
	Name string `xml:"name,attr"`
	// Color is the color used to display the objects in this group
	Color string `xml:"color,attr"`
	// X is the x coordinate of the object group in tiles
	X int `xml:"x,attr"`
	// Y is the y coordinate of the object group in tiles
	Y int `xml:"y,attr"`
	// Width is the width of the object group in tiles. Meaningless.
	Width int `xml:"width,attr"`
	// Opacity is the opacity of the layer from 0 to 1.
	Opacity float64 `xml:"opacity,attr"`
	// Visible is whether the layer is shown (1) or hidden (0).
	Visible int `xml:"visible,attr"`
	// OffsetX is the rendering x offset for this object group in pixels.
	OffsetX float64 `xml:"offsetx,attr"`
	// OffsetY is the rendering y offset for this object group in pixels.
	OffsetY float64 `xml:"offsety,attr"`
	// DrawOrder is whether the objects are drawn according to the order of
	// appearance ("index") or sorted by their y-coordinate ("topdown")
	DrawOrder string `xml:"draworder,attr"`
	// Properties are the properties of the object layer
	Properties []Property `xml:"properties>property"`
	// Objects are the objects in the object layer
	Objects []Object `xml:"object"`
}

// Object is used to add custom information to a map, such as a spawn point
type Object struct {
	// ID is the unique id of the object
	ID uint32 `xml:"id,attr"`
	// Name is the name of the object
	Name string `xml:"name,attr"`
	// Type is the type of the object
	Type string `xml:"type,attr"`
	// X is the x coordinate of the object in pixels
	X float64 `xml:"x,attr"`
	// Y is the y coordinate of the object in pixels
	Y float64 `xml:"y,attr"`
	// Width is the width of the object in pixels
	Width float64 `xml:"width,attr"`
	// Height is the height of the object in pixels
	Height float64 `xml:"height,attr"`
	// Rotation is the rotation of the object in degrees
	Rotation float64 `xml:"rotation,attr"`
	// GID is a reference to the tile
	GID uint32 `xml:"gid,attr"`
	// Visible is whether the object is shown (1) or hidden (0)
	Visible int `xml:"visible,attr"`
	// Template is a reference to a template file
	Template string `xml:"template,attr"`
	// Properties are the properties of the object
	Properties []Property `xml:"properties>property"`
	// Ellipses are any elliptical shapes
	Ellipses []Ellipse `xml:"ellipse"`
	// Polygons are any polygon shapes
	Polygons []Polygon `xml:"polygon"`
	// Polylines are any poly line shapes
	Polylines []Polyline `xml:"polyline"`
	// Text is any text
	Text []Text `xml:"text"`
	// Images is any image in the object
	Images []Image `xml:"image"`
}

// Ellipse is an elliptical shape
type Ellipse struct{}

// Point is a single point located at the object's position
type Point struct{}

// Polygon is a polygon shape
type Polygon struct {
	// Points are a list of x,y coordinates in pixels
	Points string `xml:"points,attr"`
}

// Polyline is a polygon that doesn't have to close
type Polyline struct {
	// Points are a list of x,y coordinates in pixels
	Points string `xml:"points,attr"`
}

// Text is a text object
type Text struct {
	// FontFamily is the font family used
	FontFamily string `xml:"fontfamily,attr"`
	// PixelSize is the size of the font in pixels
	PixelSize float64 `xml:"pixelsize,attr"`
	// Wrap is whether word wrapping is enabled (1) or disabled (0)
	Wrap int `xml:"wrap,attr"`
	// Color is the color of the text in #AARRGGBB or #RRGGBB format
	Color string `xml:"color,attr"`
	// Bold is whether the font is bold (1) or not (0)
	Bold int `xml:"bold,attr"`
	// Italic is whether the font is italic (1) or not (0)
	Italic int `xml:"italic,attr"`
	// Underline is whether a line should be drawn below the text (1) or not (0)
	Underline int `xml:"underline,attr"`
	// Strikeout is whether a line should be drawn through the text (1) or not (0)
	Strikeout int `xml:"strikeout,attr"`
	// Kerning is whether kerning should be used while rendering the text (1) or
	// not (0)
	Kerning int `xml:"kerning,attr"`
	// Halign is the horizontal allignment of the text within the object (left,
	// center or right)
	Halign string `xml:"halign,attr"`
	// Valign is the vertical allignment of the text within the object (top,
	// center or bottom)
	Valign string `xml:"valign,attr"`
	// CharData is the character data of the text element
	CharData string `xml:",chardata"`
}

// ImageLayer is a tile layer that contains a reference to an image
type ImageLayer struct {
	// Name is the name of the image layer
	Name string `xml:"name,attr"`
	// OffsetX is the rendering x offset of the image layer in pixels
	OffsetX float64 `xml:"offsetx,attr"`
	// OffsetY is the rendering y offset of the image layer in pixels
	OffsetY float64 `xml:"offsety,attr"`
	// X is the x position of the image layer in pixels
	X float64 `xml:"x,attr"`
	// Y is the y position of the image layer in pixels
	Y float64 `xml:"y,attr"`
	// Opacity is the opacity of the layer from 0 to 1
	Opacity float64 `xml:"opacity,attr"`
	// Visibile indicates whether the layer is shown (1) or hidden (0)
	Visible int `xml:"visible,attr"`
	// Properties are the properties of the layer
	Properties []Property `xml:"properties>property"`
	// Images are the images of the layer
	Images []Image `xml:"image"`
}

// Group is a root element to organize the layers
type Group struct {
	// Name is the name of the group layer
	Name string `xml:"name,attr"`
	// OffsetX is the x offset of the group layer in pixels
	OffsetX float64 `xml:"offsetx,attr"`
	// OffsetY is the y offset of the group layer in pixels
	OffsetY float64 `xml:"offsety,attr"`
	// Opacity is the opacity of the layer from 0 to 1
	Opacity float64 `xml:"opacity,attr"`
	// Visible is whether the layer is shown (1) or hidden (0)
	Visible int `xml:"visible,attr"`
	// Properties are the properties of the group
	Properties []Property `xml:"properties>property"`
	// Layers are the layers of the group
	Layers []Layer `xml:"layer"`
	// ObjectGroups are the object groups of the group
	ObjectGroups []ObjectGroup `xml:"objectgroup"`
	// ImageLayers are the image layers of the group
	ImageLayers []ImageLayer `xml:"imagelayer"`
	// Groups are the child groups in the group
	Group []Group `xml:"group"`
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (o *ObjectGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type objectGroup ObjectGroup
	og := objectGroup{
		Opacity:   1,
		Visible:   1,
		DrawOrder: "topdown",
	}
	if err := d.DecodeElement(&og, &start); err != nil {
		return err
	}
	*o = (ObjectGroup)(og)
	return nil
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (o *Object) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type object Object
	obj := object{
		Visible: 1,
	}
	if err := d.DecodeElement(&obj, &start); err != nil {
		return err
	}
	*o = (Object)(obj)
	if o.Template != "" {
		tmpl := Template{}
		f, err := os.Open(path.Join(path.Dir(TMXURL), o.Template))
		defer f.Close()
		if err != nil {
			return err
		}
		b, _ := ioutil.ReadAll(f)
		err = xml.Unmarshal(b, &tmpl)
		if err != nil {
			return err
		}
		if len(o.Ellipses) == 0 {
			o.Ellipses = tmpl.Objects[0].Ellipses
		}
		if o.GID == 0 {
			o.GID = tmpl.Objects[0].GID
		}
		if o.Height == 0 {
			o.Height = tmpl.Objects[0].Height
		}
		if len(o.Images) == 0 {
			o.Images = tmpl.Objects[0].Images
		}
		if o.Name == "" {
			o.Name = tmpl.Objects[0].Name
		}
		if len(o.Polygons) == 0 {
			o.Polygons = tmpl.Objects[0].Polygons
		}
		if len(o.Polylines) == 0 {
			o.Polylines = tmpl.Objects[0].Polylines
		}
		if len(o.Properties) == 0 {
			o.Properties = tmpl.Objects[0].Properties
		}
		if o.Rotation == 0 {
			o.Rotation = tmpl.Objects[0].Rotation
		}
		if len(o.Text) == 0 {
			o.Text = tmpl.Objects[0].Text
		}
		if o.Type == "" {
			o.Type = tmpl.Objects[0].Type
		}
		if o.Visible == 1 {
			o.Visible = tmpl.Objects[0].Visible
		}
		if o.Width == 0 {
			o.Width = tmpl.Objects[0].Width
		}
		if o.X == 0 {
			o.X = tmpl.Objects[0].X
		}
		if o.Y == 0 {
			o.Y = tmpl.Objects[0].Y
		}
	}
	return nil
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (t *Text) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type text Text
	txt := text{
		FontFamily: "sans-serif",
		PixelSize:  16,
		Wrap:       0,
		Color:      "#000000",
		Bold:       0,
		Italic:     0,
		Strikeout:  0,
		Underline:  0,
		Kerning:    1,
		Halign:     "left",
		Valign:     "top",
	}
	if err := d.DecodeElement(&txt, &start); err != nil {
		return err
	}
	*t = (Text)(txt)
	return nil
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (i *ImageLayer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type imageLayer ImageLayer
	il := imageLayer{
		Opacity: 1,
		Visible: 1,
	}
	if err := d.DecodeElement(&il, &start); err != nil {
		return err
	}
	*i = (ImageLayer)(il)
	return nil
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (g *Group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type group Group
	gr := group{
		Opacity: 1,
		Visible: 1,
	}
	if err := d.DecodeElement(&gr, &start); err != nil {
		return err
	}
	*g = (Group)(gr)
	return nil
}
