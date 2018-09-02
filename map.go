package tmx

import "encoding/xml"

// Map is the root element of a TMX map
type Map struct {
	// Version is the TMX format version
	Version string `xml:"version,attr"`
	// TiledVersion is the Version of Tiled Map Editor used to generate the TMX
	TiledVersion string `xml:"tiledversion,attr"`
	// Orientation is the orientation of the map. Tiled supports “orthogonal”,
	// “isometric”, “staggered” and “hexagonal”
	Orientation string `xml:"orientation,attr"`
	// RenderOrder is The order in which tiles on tile layers are rendered.
	// Valid values are right-down (the default), right-up, left-down and left-up.
	// In all cases, the map is drawn row-by-row.
	// (only supported for orthogonal maps at the moment)
	RenderOrder string `xml:"renderorder,attr"`
	// Width is the map width in tiles
	Width int `xml:"width,attr"`
	// Height is the map height in tiles
	Height int `xml:"height,attr"`
	// TileWidth is the width of each tile in pixels
	TileWidth int `xml:"tilewidth,attr"`
	// TileHeight is the height of each tile in pixels
	TileHeight int `xml:"tileheight,attr"`
	// HexSideLength determines the width or height (depending on the staggered
	// axis) of the tile’s edge, in pixels. Only for hexagonal maps.
	HexSideLength int `xml:"hexsidelength,attr"`
	// StaggerAxis is for staggered and hexagonal maps, determines which axis
	// (“x” or “y”) is staggered.
	StaggerAxis string `xml:"staggeraxis,attr"`
	// StaggerIndex is for staggered and hexagonal maps, determines whether the
	// “even” or “odd” indexes along the staggered axis are shifted.
	StaggerIndex string `xml:"staggerindex,attr"`
	// BackgroundColor is the background color of the map. Is of the form #AARRGGBB
	BackgroundColor string `xml:"backgroundcolor,attr"`
	// NextObjectID stores the next object id available for new objects.
	NextObjectID int `xml:"nextobjectid,attr"`
	// Properties are the properties of the map
	Properties []Property `xml:"properties>property"`
	// Tilesets are the tilesets of the map
	Tilesets []Tileset `xml:"tileset"`
	// Layers are the layers of the map
	Layers []Layer `xml:"layer"`
	// ObjectGroups are the object groups of the map
	ObjectGroups []ObjectGroup `xml:"objectgroup"`
	// ImageLayers are the image layers of the map
	ImageLayers []ImageLayer `xml:"imagelayer"`
	// Groups are the groups of the map
	Groups []Group `xml:"group"`
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (m *Map) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type maph Map
	ma := maph{
		RenderOrder: "right-down",
	}
	if err := d.DecodeElement(&ma, &start); err != nil {
		return err
	}
	*m = (Map)(ma)
	return nil
}
