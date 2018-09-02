package tmx

// Template is a separate file that contains the template root element, a
// map object and a tileset element that points to an external tileset if
// the object is a tile object
type Template struct {
	// Tilesets are the tilesets for the template
	Tilesets []Tileset `xml:"tileset"`
	// Objects are the template objects
	Objects []Object `xml:"object"`
}
