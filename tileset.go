package game_manager

type Tileset struct {
	XMLName    xml.Name `xml:"tileset"`
	name       string   `xml:"name"`
	firstGidID int      `xml:"firstgid"`
	tileWidth  int      `xml:"tilewidth"`
	tileHeight int      `xml:"tileheight"`
	spacing    int      `xml:"spacing"`
	margin     int      `xml:"margin"`
	numColumns int
}

type TilesetImage struct {
	source string `xml:"source"`
	width  int    `xml:"width"`
	height int    `xml:"height"`
}
