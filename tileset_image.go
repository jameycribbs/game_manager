package game_manager

type TilesetImage struct {
	XMLName xml.Name `xml:"image"`
	source  string   `xml:"source"`
	width   int      `xml:"width"`
	height  int      `xml:"height"`
}
