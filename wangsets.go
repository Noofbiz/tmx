package tmx

// WangSet is a wang set from the TMX map
type WangSet struct {
	// Name is the name of the wang set
	Name string `xml:"name,attr"`
	// Tile is the local tile id of the wang set
	Tile uint32 `xml:"id,attr"`
	// WangCornerColor is a color that can be used to define the corner of a
	// Wang tile
	WangCornerColors []WangCornerColor `xml:"wangcornercolor"`
	// WangEdgeColor is a color that can be used to define the edge of a Wang tile
	WangEdgeColors []WangEdgeColor `xml:"wangedgecolor"`
	// WangTile defines a Wang Tile
	WangTiles []WangTile `xml:"wangtile"`
}

// WangCornerColor is a color that can be used to define the edge of a Wang
// tile
type WangCornerColor struct {
	// Name is the name of this color
	Name string `xml:"name,attr"`
	// Color is the color in #RRGGBB format
	Color string `xml:"color,attr"`
	// Tile is the tile ID of the tile representing this color
	Tile uint32 `xml:"tile,attr"`
	// Probability is the relative probability that this color is chosen
	Probability float64 `xml:"probability,attr"`
}

// WangEdgeColor is a color that can be used to define the edge of a Wang tile
type WangEdgeColor struct {
	// Name is the name of this color
	Name string `xml:"name,attr"`
	// Color is the color in #RRGGBB format
	Color string `xml:"color,attr"`
	// Tile is the tile ID of the tile representing this color
	Tile uint32 `xml:"tile,attr"`
	// Probability is the relative probability that this color is chosen
	Probability float64 `xml:"probability,attr"`
}

// WangTile is a wang tile. It refers to a tile in the tileset and associates
// it with a Wang ID
type WangTile struct {
	// TileID is the tile ID
	TileID uint32 `xml:"tileid,attr"`
	// WangID is a 32-bit unsigned integer stored in the format 0xCECECECE where
	// C is a corner color and each E is an edge color, from right to left
	// clockwise, starting with the top edge.
	// It's stored as a string here rather than a uint32 it's easer for the end
	// user of the library to read / parse themselves.
	WangID string `xml:"wangid,attr"`
}
