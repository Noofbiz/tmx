package tmx

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
)

// Tileset is a tileset used for the map
type Tileset struct {
	// FirstGID is  the first global tile ID of this tileset (this global ID maps
	// to the first tile in this tileset)
	FirstGID uint32 `xml:"firstgid,attr"`
	// Source is the location of the external tilemap TSX file, if any
	Source string `xml:"source,attr"`
	// Name is the name of the tileset
	Name string `xml:"name,attr"`
	// TileWidth is the (maximum) width of tiles in the tileset
	TileWidth int `xml:"tilewidth,attr"`
	// TileHeight is the (maximum) height of the tiles in the tileset
	TileHeight int `xml:"tileheight,attr"`
	// Spacing is the spacing of the tiles in pixels between the tiles in the tileset
	Spacing int `xml:"spacing,attr"`
	// Margin is the margin around the tiles in pixels of the tiles in the tileset
	Margin float64 `xml:"margin,attr"`
	// TileCount is the number of tiles in the tileset
	TileCount int `xml:"tilecount,attr"`
	// Columns is the number of tile columns in the tileset
	Columns int `xml:"columns,attr"`
	// TileOffset is used to specify an offset in pixels, to be applied when
	// drawing a tile from the related tileset. When not present, no offset
	// is applied
	TileOffset []TileOffset `xml:"tileoffset"`
	// Grid is is only used in case of isometric orientation, and determines how
	// tile overlays for terrain and collision information are rendered
	Grid []Grid `xml:"grid"`
	// Properties are the custom properties of the tileset
	Properties []Property `xml:"properties>property"`
	// Image is the image associated with the tileset
	Image []Image `xml:"image"`
	// TerrainTypes are the terraintypes associated with the tileset
	TerrainTypes []Terrain `xml:"terraintypes>terrain"`
	// Tiles are tiles in the tileset
	Tiles []Tile `xml:"tile"`
	// WangSets contain the list of wang sets defined for this tileset
	WangSets []WangSet `xml:"wangsets>wangset"`
}

// TileOffset is used to specify an offset in pixels, to be applied when
// drawing a tile from the related tileset. When not present, no offset is
// applied.
type TileOffset struct {
	// X is the horizontal offset in pixels
	X float64 `xml:"x,attr"`
	// Y is the vertical offset in pixels. Positive is down.
	Y float64 `xml:"y,attr"`
}

// Grid is only used in case of isometric orientation, and determines
// how tile overlays for terrain and collision information are rendered.
type Grid struct {
	//  Orientation of the grid for the tiles in this tileset (orthogonal or
	// isometric)
	Orientation string `xml:"orientation,attr"`
	// Width is the width of a grid cell
	Width float64 `xml:"width,attr"`
	// Height is the height of a grid cell
	Height float64 `xml:"height,attr"`
}

// Image is data for an image file
type Image struct {
	// Format is used for embedded images, in combination with a data child element.
	// Valid values are file extensions like png, gif, jpg, bmp, etc.
	Format string `xml:"format,attr"`
	// Source is the reference to the tileset image file.
	Source string `xml:"source,attr"`
	// Transparent defines a specific color that is treated as transparent (example
	// value: “#FF00FF” for magenta). Up until Tiled 0.12, this value is written
	// out without a # but this is planned to change.
	Transparent string `xml:"trans,attr"`
	// Width is the image width in pixels
	Width float64 `xml:"width,attr"`
	// Height is the image height in pixels
	Height float64 `xml:"height,attr"`
	// Data is the image data
	Data []Data `xml:"data"`
}

// Terrain is a terrain
type Terrain struct {
	// Name is the name of the terrain type
	Name string `xml:"name,attr"`
	// Tile is the id of the local tile that represents the terrain
	Tile uint32 `xml:"tile,attr"`
}

// Tile is a single tile in a tile set
type Tile struct {
	// ID is the local tile id within its tileset
	ID uint32 `xml:"id,attr"`
	// Type is the type of the tile
	Type string `xml:"type,attr"`
	// Terrain defines the terrain type of each corner of the tile, given as
	// comma-separated indexes in the terrain types array in the order top-left,
	// top-right, bottom-left, bottom-right. Leaving out a value means that corner
	// has no terrain. (optional)
	Terrain string `xml:"terrain,attr"`
	// Probability is a percentage indicating the probability that this tile is
	// chosen when it competes with others while editing with the terrain tool.
	Probability float64 `xml:"probability,attr"`
	// Properties are the custom properties of the tile
	Properties []Property `xml:"properties>property"`
	// Image is the image associated with the tile
	Image []Image `xml:"image"`
	// ObjectGroups are a group of objects
	ObjectGroup []ObjectGroup `xml:"objectgroup"`
	// AnimationFrames are any frames to animate the tile
	AnimationFrames []Frame `xml:"animation>frame"`
}

// Frame is an animation frame
type Frame struct {
	// TileID is the local id of a tile in the tileset
	TileID uint32 `xml:"tileid,attr"`
	// Duration is how long (in milliseconds) this frame should be displayed
	// before advancing to the next frame
	Duration float64 `xml:"duration,attr"`
}

// UnmarshalXML implements the encoding/xml Unmarshaler interface
func (t *Tileset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tileset Tileset
	ts := tileset{}
	if err := d.DecodeElement(&ts, &start); err != nil {
		return err
	}
	*t = (Tileset)(ts)
	if t.Source != "" {
		t2 := Tileset{}
		f, err := os.Open(path.Join(path.Dir(TMXURL), t.Source))
		defer f.Close()
		if err != nil {
			return err
		}
		b, _ := ioutil.ReadAll(f)
		err = xml.Unmarshal(b, &t2)
		if err != nil {
			return err
		}
		t.Name = t2.Name
		t.TileWidth = t2.TileWidth
		t.TileHeight = t2.TileHeight
		t.Spacing = t2.Spacing
		t.Margin = t2.Margin
		t.TileCount = t2.TileCount
		t.Columns = t2.Columns
		t.TileOffset = t2.TileOffset
		t.Grid = t2.Grid
		t.Image = t2.Image
		t.Properties = t2.Properties
		t.TerrainTypes = t2.TerrainTypes
		t.Tiles = t2.Tiles
		t.WangSets = t2.WangSets
	}
	return nil
}
